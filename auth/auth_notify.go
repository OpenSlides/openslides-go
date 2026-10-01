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
	logoutMU                sync.RWMutex
)

type BlockedSession struct {
	ID        int
	SessionID string
	Timestamp *time.Time
}

type LogoutListener struct {
	blockedSessions map[string]int
	lookup          environment.Environmenter
	pool            *pgxpool.Pool
}

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

func (ll *LogoutListener) Close() {
	ll.pool.Close()
}

func (ll *LogoutListener) IsBlocked(sessionID string) bool {
	logoutMU.RLock()
	_, found := ll.blockedSessions[sessionID]
	logoutMU.RUnlock()
	return found
}

func (ll *LogoutListener) populate(ctx context.Context) error {
	logoutMU.Lock()
	// Clear current map
	clear(ll.blockedSessions)

	// Fetch all
	rows, err := ll.pool.Query(ctx, "SELECT * FROM blocked_sessions_t;")
	if err != nil {
		logoutMU.Unlock()
		return &authError{"err fetch all in blocked session handler", err}
	}
	defer func() {
		rows.Close()
		logoutMU.Unlock()
	}()

	for rows.Next() {
		var b BlockedSession
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
	return nil
}

// listenOnLogouts listen on logout events and closes the connections.
func (ll *LogoutListener) listenOnLogouts(ctx context.Context, errHandler func(error)) {
	defer ll.Close()

	if errHandler == nil {
		errHandler = func(error) {}
	}

	var err error
	for {
		select {
		// case <-ctx.Err():
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				return
			}
			// Datastore Fetch
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
