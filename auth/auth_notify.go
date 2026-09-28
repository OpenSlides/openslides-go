package auth

import (
	"context"
	"fmt"
	"strings"

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

type LogoutListener struct {
	blockedSessions map[string]int
	lookup          environment.Environmenter
	pool            *pgxpool.Pool
}

func NewLogoutListener(lookup environment.Environmenter) (*LogoutListener, error) {
	// Create Postgres Pool
	addr, err := postgresDSN(lookup)
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Create Logout Listener
	l := &LogoutListener{
		blockedSessions: make(map[string]int),
		lookup:          lookup,
		pool:            pool,
	}

	// Populate blocklist with database content at start-up
	l.populate()

	// Regularily populate blocklist

	return l, nil
}

func (l *LogoutListener) IsBlocked(sessionID string) bool {
	return true
}

func (l *LogoutListener) populate() {
	// Clear current map
	clear(l.blockedSessions)

	// Fetch all
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
