package pgtest

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest/v3"
)

//go:generate go run ./copy_sql

//go:embed sql/schema_relational.sql
var schemaSQL string

//go:embed sql/base_data.sql
var baseDataSQL string

type PostgresTest struct {
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource

	Env map[string]string

	pgxConfig *pgx.ConnConfig
}

// NewPostgresTest creates a PostgresTest instance to test against a postgres
// server in a docker container.
func NewPostgresTest(ctx context.Context) (tp_ *PostgresTest, err error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("connect to docker: %w", err)
	}

	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=openslides",
			"POSTGRES_DB=database",
		},
	}

	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		return nil, fmt.Errorf("start postgres container: %w", err)
	}

	port := resource.GetPort("5432/tcp")
	addr := fmt.Sprintf(`user=postgres password='openslides' host=localhost port=%s dbname=database`, port)
	config, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	tp := &PostgresTest{
		dockerPool:     pool,
		dockerResource: resource,
		pgxConfig:      config,

		Env: map[string]string{
			"DATABASE_HOST": "localhost",
			"DATABASE_PORT": port,
			"DATABASE_NAME": "database",
			"DATABASE_USER": "postgres",
		},
	}

	defer func() {
		if err != nil {
			if err := tp.Close(); err != nil {
				log.Println("Closing postgres: %w", err)
			}
		}
	}()

	if err := tp.addSchema(ctx); err != nil {
		return nil, fmt.Errorf("add schema: %w", err)
	}

	return tp, nil
}

func (tp *PostgresTest) Close() error {
	if err := tp.dockerPool.Purge(tp.dockerResource); err != nil {
		return fmt.Errorf("purge postgres container: %w", err)
	}
	return nil
}

// Conn returns a pgx connection to the postgres server.
func (tp *PostgresTest) Conn(ctx context.Context) (*pgx.Conn, error) {
	var conn *pgx.Conn

	for {
		var err error
		conn, err = pgx.ConnectConfig(ctx, tp.pgxConfig)
		if err == nil {
			return conn, nil
		}

		select {
		case <-time.After(200 * time.Millisecond):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (tp *PostgresTest) addSchema(ctx context.Context) error {
	conn, err := tp.Conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, schemaSQL); err != nil {
		return fmt.Errorf("adding schema: %w", PrityPostgresError(err, schemaSQL))
	}
	return nil
}

// Cleanup uses t.cleanup to register a cleanup function after the test is done.
//
// This can be used to reuse a PostgresTest without restarting the postgres
// server.
func (tp *PostgresTest) Cleanup(t *testing.T) {
	t.Cleanup(func() {
		if err := tp.reset(t.Context()); err != nil {
			t.Logf("Cleanup database: %v", err)
		}
	})
}

func (tp *PostgresTest) reset(ctx context.Context) error {
	// Use different database for drop database
	config := tp.pgxConfig.Copy()
	config.Database = "postgres"
	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("create connection: %w", err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, `DROP DATABASE database WITH (FORCE);`); err != nil {
		return fmt.Errorf("dropping database: %w", err)
	}

	if _, err := conn.Exec(ctx, `CREATE DATABASE database;`); err != nil {
		return fmt.Errorf("recreating database: %w", err)
	}

	if err := tp.addSchema(ctx); err != nil {
		return fmt.Errorf("adding schema: %w", err)
	}

	return nil
}

func (tp *PostgresTest) addBaseData(ctx context.Context) error {
	conn, err := tp.Conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}

	if _, err := conn.Exec(ctx, baseDataSQL); err != nil {
		return fmt.Errorf("adding example data: %w", PrityPostgresError(err, baseDataSQL))
	}

	return nil
}

func PrityPostgresError(err error, sql string) error {
	var errPG *pgconn.PgError
	if errors.As(err, &errPG) {
		line := getLineNumber(sql, int(errPG.Position))
		contextStart := max(0, int(errPG.Position)-100)
		contextEnd := min(len(sql), int(errPG.Position)+100)
		context := sql[contextStart:contextEnd]

		return fmt.Errorf("PostgreSQL error at line %d, byte position %d: %s\nContext: ...%s...",
			line, int(errPG.Position), errPG.Message, context)
	}
	return err
}

func getLineNumber(text string, bytePos int) int {
	if bytePos > len(text) {
		bytePos = len(text)
	}

	line := 1
	for i := range bytePos {
		if text[i] == '\n' {
			line++
		}
	}
	return line
}
