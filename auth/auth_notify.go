package auth

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	envPostgresHost         = environment.NewVariable("DATABASE_HOST", "localhost", "Postgres Host.")
	envPostgresPort         = environment.NewVariable("DATABASE_PORT", "5432", "Postgres Post.")
	envPostgresDatabase     = environment.NewVariable("DATABASE_NAME", "openslides", "Postgres User.")
	envPostgresUser         = environment.NewVariable("DATABASE_USER", "openslides", "Postgres Database.")
	envPostgresPasswordFile = environment.NewVariable("DATABASE_PASSWORD_FILE", "/run/secrets/postgres_password", "Postgres Password.")
)

// Datastructure for a blocked session. Wraps blocked session ID and when it has been blocked
type blockedSession struct {
	ID        int
	SessionID string
	Timestamp *time.Time
}

// Listens to logouts (in the form of blocked sessions) in the database. In that event it emits a signal to all listening goroutines
type LogoutListener struct {
	blockedSessions map[string]int
	lookup          environment.Environmenter
	pool            *pgxpool.Pool
	notifyChannel   *sync.Cond
	notifyMU        sync.Mutex
}

// Creates a new LogoutListener. Opens up a pool which is closed automatically once the background function is called
func NewLogoutListener(lookup environment.Environmenter) (*LogoutListener, func(context.Context, func(error)), error) {
	// Create Postgres Pool
	addr, err := postgresDSN(lookup)
	if err != nil {
		return nil, nil, err
	}

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, nil, err
	}

	// Close pool if an error occurs prior to returning background func
	defer func() {
		if r := recover(); r != nil {
			pool.Close()
		}
	}()

	// Create Logout Listener
	ll := &LogoutListener{
		blockedSessions: make(map[string]int),
		lookup:          lookup,
		pool:            pool,
	}

	ll.notifyChannel = sync.NewCond(&ll.notifyMU)

	// Regularily populate blocklist as background task
	background := func(ctx context.Context, errorHandler func(error)) {
		// Populate blocklist with database content at start-up
		err := ll.populate(ctx)
		if err != nil {
			errorHandler(err)
			return
		}

		go ll.listenOnLogouts(ctx, errorHandler)
	}

	return ll, background, nil
}

// Closes the pgx pool used to listen to blocked session events
func (ll *LogoutListener) Close() {
	ll.pool.Close()
}

// Returns true, if the given session ID is present in the blocked sessions database table
func (ll *LogoutListener) IsBlocked(sessionID string) bool {
	_, found := ll.blockedSessions[sessionID]
	return found
}

// Emits signal to all goroutines listening for notifies on the blocked sessions database table
func (ll *LogoutListener) NotifyCond() *sync.Cond {
	return ll.notifyChannel
}

func (ll *LogoutListener) populate(ctx context.Context) error {
	ll.notifyMU.Lock()
	defer ll.notifyMU.Unlock()
	// Clear current map
	clear(ll.blockedSessions)

	// Fetch all
	rows, err := ll.pool.Query(ctx, "SELECT * FROM blocked_sessions_t;")
	if err != nil {
		return &authError{"err fetch all in blocked session handler", err}
	}
	defer rows.Close()

	for rows.Next() {
		var b blockedSession
		err := rows.Scan(&b.ID, &b.SessionID) //, &b.Timestamp)

		if err != nil {
			return &authError{"err scaning row in blocked session handler", err}
		}

		if b.Timestamp == nil || time.Since(*b.Timestamp) <= 30*time.Minute {
			ll.blockedSessions[b.SessionID] = b.ID
		} else {
			fmt.Printf("Session ID has been blocked but is outdated: %v", b.SessionID)
		}
	}

	ll.NotifyCond().Broadcast()

	return nil
}

// listenOnLogouts listen on logout events and closes the connections.
func (ll *LogoutListener) listenOnLogouts(ctx context.Context, errHandler func(error)) {
	defer ll.Close()

	if errHandler == nil {
		errHandler = func(error) {}
	}

	// Connect to blockes sessions notify triggers
	conn, err := ll.pool.Acquire(ctx)
	if err != nil {
		errHandler(fmt.Errorf("acquiring connection for LISTEN: %w", err))
		return
	}
	defer conn.Release()

	if _, err = conn.Exec(ctx, "LISTEN blocked_sessions_notify"); err != nil {
		errHandler(fmt.Errorf("subscribing to LISTEN: %w", err))
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				return
			}

			_, err := conn.Conn().WaitForNotification(ctx)
			if err != nil {
				if err == context.Canceled {
					return
				}
				errHandler(fmt.Errorf("waiting for blocked session notification: %w", err))
				return
			}

			// Populate blocklist
			err = ll.populate(ctx)

			if err != nil {
				errHandler(err)
				return
			}

			time.Sleep(time.Second)
		}
	}
}

// TODO: This is the same as in flow_postgres.go. Should be reused
// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

func postgresDSN(lookup environment.Environmenter) (string, error) {
	password, err := environment.ReadSecret(lookup, envPostgresPasswordFile)
	if err != nil {
		return "", fmt.Errorf("reading postgres password: %w", err)
	}

	return postgresConfigString(
		encodePostgresConfig(envPostgresHost.Value(lookup)),
		encodePostgresConfig(envPostgresPort.Value(lookup)),
		encodePostgresConfig(envPostgresDatabase.Value(lookup)),
		encodePostgresConfig(envPostgresUser.Value(lookup)),
		encodePostgresConfig(password),
	), nil
}

func postgresConfigString(host, port, db, user, password string) string {
	return fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(user),
		encodePostgresConfig(password),
		encodePostgresConfig(host),
		encodePostgresConfig(port),
		encodePostgresConfig(db),
	)
}
