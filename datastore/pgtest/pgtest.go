package pgtest

import (
	"context"
	"crypto/sha256"
	_ "embed" // Needed for embedding
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ory/dockertest/v4"
)

// Go embed can only include files inside the same directory or sub directories.
// But for security reasons, it can not include files outside the directory.
// Therefore it is necessary to copry the files to create the sql-schema to this
// folder. The embedding is necessary, to other repositories like the
// vote-service, do not need to include the meta-repo as a sub repo.

//go:generate go run ./copy_sql

//go:embed sql/schema_relational.sql
var schemaSQL string

//go:embed sql/base_data.sql
var baseDataSQL string

var pool dockertest.ClosablePool

// RunTests has to be called in TestMain by every package, that uses
// NewPostgresTest.
func RunTests(m *testing.M) int {
	ctx := context.Background()
	var err error
	pool, err = dockertest.NewPool(ctx, "",
		dockertest.WithMaxWait(2*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	code := m.Run()
	// Close removes all tracked containers/networks and closes the client.
	// Call before os.Exit — deferred functions do not run after os.Exit.
	pool.Close(ctx)
	return code
}

// PostgresTest is a test helper for postgres.
//
// It creates or reuses a postgres container. Each test gets its own database.
// After all test complete, the container ist shut down.
//
// When using this in subtests, either create a new PostgresTest for each
// subtest or make sure, that all subtest do not conflict with each other, since
// they are using the same database.
//
// You have to use the following code in each test-package, that wants to use this structure.
//
//	func TestMain(m *testing.M) {
//		os.Exit(pgtest.RunTests(m))
//	}
type PostgresTest struct {
	Env map[string]string

	pgxConfig *pgx.ConnConfig
}

// NewPostgresTest creates a PostgresTest instance to test against a postgres
// server in a docker container.
func NewPostgresTest(t *testing.T) (*PostgresTest, error) {
	ctx := t.Context()
	postgres, err := pool.Run(ctx, "postgres",
		dockertest.WithTag("17"),
		dockertest.WithEnv([]string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=openslides",
			"POSTGRES_DB=postgres",
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("run postgres: %w", err)
	}

	database := dbName(t)
	port := postgres.GetPort("5432/tcp")

	adminAddr := fmt.Sprintf(`user=postgres password='openslides' host=localhost port=%s dbname=postgres`, port)
	adminConfig, err := pgx.ParseConfig(adminAddr)
	if err != nil {
		return nil, fmt.Errorf("parse admin config: %w", err)
	}

	err = pool.Retry(ctx, 30*time.Second, func() error {
		conn, err := pgx.ConnectConfig(ctx, adminConfig)
		if err != nil {
			return fmt.Errorf("create admin connection: %w", err)
		}
		defer conn.Close(ctx)

		if _, err = conn.Exec(ctx, "CREATE DATABASE "+database); err != nil {
			return fmt.Errorf("send create database command: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create test database: %w", err)
	}

	addr := fmt.Sprintf(`user=postgres password='openslides' host=localhost port=%s dbname=%s`, port, database)
	config, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	tp := &PostgresTest{
		pgxConfig: config,

		Env: map[string]string{
			"DATABASE_HOST": "localhost",
			"DATABASE_PORT": port,
			"DATABASE_NAME": database,
			"DATABASE_USER": "postgres",
		},
	}

	if err := tp.addSchema(ctx); err != nil {
		return nil, fmt.Errorf("add schema: %w", err)
	}

	if err := tp.addBaseData(ctx); err != nil {
		return nil, fmt.Errorf("add base data: %w", err)
	}

	return tp, nil
}

func dbName(t *testing.T) string {
	name := "test_" + t.Name()

	name = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r + 32
		}
		return '_'
	}, name)

	if len(name) > 63 {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(name)))
		name = name[:47] + "_" + hash[:15]
	}

	return name
}

// Conn returns a pgx connection to the postgres server.
func (tp *PostgresTest) Conn(ctx context.Context) (*pgx.Conn, error) {
	var conn *pgx.Conn
	var err error
	for range 100 {
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
	return nil, fmt.Errorf("getting connections 100 times: %w", err)
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

// AddData adds data in yaml-format to the database.
func (tp *PostgresTest) AddData(ctx context.Context, data string) error {
	conn, err := tp.Conn(ctx)
	if err != nil {
		return fmt.Errorf("open connection: %w", err)
	}
	defer conn.Close(ctx)

	if err := insertTestData(ctx, conn, data); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}

	return nil
}

// Flow returns a flow that is using the postgres instance.
func (tp *PostgresTest) Flow(ctx context.Context) (*datastore.FlowPostgres, error) {
	flow, init, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
	if err != nil {
		return nil, fmt.Errorf("create postgres flow: %w", err)
	}
	if err := init(ctx); err != nil {
		return nil, fmt.Errorf("init postgres: %w", err)
	}
	return flow, nil
}

// PrityPostgresError returns a formatted error message for PostgreSQL errors.
func PrityPostgresError(err error, sql string) error {
	var errPG *pgconn.PgError
	if errors.As(err, &errPG) {
		line := getLineNumber(sql, int(errPG.Position))
		contextStart := max(0, int(errPG.Position)-100)
		contextEnd := min(len(sql), int(errPG.Position)+100)
		context := sql[contextStart:contextEnd]

		return fmt.Errorf("postgreSQL error at line %d, byte position %d: %s\nContext: %s",
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
