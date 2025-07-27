package datastore_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
)

func TestFlowPostgres(t *testing.T) {
	ctx := t.Context()

	t.Parallel()
	if testing.Short() {
		t.Skip("Postgres Test")
	}

	tp, err := newTestPostgres(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	for _, tt := range []struct {
		name   string // Name of the test
		insert string
		expect map[string][]byte // expected data. Uses a get request on all keys of the expect map
	}{
		{
			"Same fqid",
			`
			INSERT INTO theme (name, accent_500, primary_500, warn_500) values ('standard theme', '#123456', '#123456', '#123456');
			`,
			map[string][]byte{
				"theme/1/name": []byte(`"standard theme"`),
			},
		},
		{
			"different fqid",
			`
			INSERT INTO theme (name, accent_500, primary_500, warn_500) values ('standard theme', '#123456', '#123456', '#123456');
			INSERT INTO "user" (id, username, first_name) values (42,'hugo', 'Hugo');
			`,
			map[string][]byte{
				"user/42/username":   []byte(`"hugo"`),
				"user/42/first_name": []byte(`"Hugo"`),
				"theme/1/name":       []byte(`"standard theme"`),
			},
		},
		{
			"Empty Data",
			``,
			map[string][]byte{
				"motion/2/title": nil,
			},
		},
		{
			"Empty Data on id list",
			`
			INSERT INTO "user" (id, username) values (10,'hugo');
			`,
			map[string][]byte{
				"user/10/meeting_user_ids": nil,
			},
		},
		{
			"Decimal",
			`
			INSERT INTO "user" (id, username, default_vote_weight) values (15,'hugo', 1.5);
			`,
			map[string][]byte{
				"user/15/default_vote_weight": []byte(`"1.500000"`),
			},
		},
		{
			"Boolean",
			`
			INSERT INTO theme (name, accent_500, primary_500, warn_500) VALUES ('standard theme', '#123456', '#123456', '#123456');
			INSERT INTO organization (id, name, default_language, theme_id, enable_electronic_voting) VALUES (1, 'my orga', 'en', 1, true);
			`,
			map[string][]byte{
				"organization/1/enable_electronic_voting": []byte(`true`),
			},
		},
		{
			"JSON",
			`
			INSERT INTO theme (name, accent_500, primary_500, warn_500) VALUES ('standard theme', '#123456', '#123456', '#123456');
			INSERT INTO organization (id, name, default_language, theme_id, saml_attr_mapping) VALUES (1, 'my orga', 'en', 1, '{"key1": "value1", "key2": "value2"}');
			`,
			map[string][]byte{
				"organization/1/saml_attr_mapping": []byte(`{"key1": "value1", "key2": "value2"}`),
			},
		},
		{
			"Text",
			`
			INSERT INTO theme (name, accent_500, primary_500, warn_500) VALUES ('standard theme', '#123456', '#123456', '#123456');
			INSERT INTO organization (id, name, default_language, theme_id, saml_private_key) VALUES (1, 'my orga', 'en', 1, 'some text');
			`,
			map[string][]byte{
				"organization/1/saml_private_key": []byte(`"some text"`),
			},
		},
		{
			"Timestamp",
			`
			INSERT INTO "user" (username, last_login) values ('hugo', '1999-01-08');
			`,
			map[string][]byte{
				"user/1/last_login": []byte(`915753600`),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			tp.cleanup(t)

			conn, err := tp.conn(ctx)
			if err != nil {
				t.Fatalf("create connection: %v", err)
			}
			defer conn.Close(ctx)

			if _, err := conn.Exec(ctx, tt.insert); err != nil {
				t.Fatalf("adding example data: %v", err)
			}

			flow, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
			if err != nil {
				t.Fatalf("NewFlowPostgres(): %v", err)
			}
			defer flow.Close()

			keys := make([]dskey.Key, 0, len(tt.expect))
			for k := range tt.expect {
				keys = append(keys, dskey.MustKey(k))
			}

			got, err := flow.Get(ctx, keys...)
			if err != nil {
				t.Fatalf("Get: %v", err)
			}

			expect := make(map[dskey.Key][]byte)
			for k, v := range tt.expect {
				expect[dskey.MustKey(k)] = v
			}

			if !reflect.DeepEqual(got, expect) {
				t.Errorf("\nGot\t\t%v\nexpect\t%v", got, expect)
			}
		})
	}
}

func TestPostgresUpdate(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
	defer cancel()

	t.Parallel()
	if testing.Short() {
		t.Skip("Postgres Test")
	}

	tp, err := newTestPostgres(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	conn, err := tp.conn(ctx)
	if err != nil {
		t.Fatalf("create connection: %v", err)
	}

	flow, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
	if err != nil {
		t.Fatalf("NewFlowPostgres(): %v", err)
	}

	keys := []dskey.Key{
		dskey.MustKey("user/1/username"),
		dskey.MustKey("theme/1/name"),
	}

	done := make(chan error)
	// TODO: When update fails, this currently blocks for ever. Maybe use a
	// timeout.
	go flow.Update(ctx, func(m map[dskey.Key][]byte, err error) {
		select {
		case <-done:
			// Only call update once.
			return
		default:
		}

		if err != nil {
			done <- fmt.Errorf("from Update callback: %w", err)
			return
		}

		if m[keys[0]] == nil {
			done <- fmt.Errorf("key %s not found", keys[0])
			return
		}

		if m[keys[1]] == nil {
			done <- fmt.Errorf("key %s not found", keys[1])
			return
		}

		done <- nil
		return
	})
	// TODO: This test could be flaky.
	time.Sleep(5 * time.Second) // TODO: How to do this without a sleep?
	sql := `
	INSERT INTO "user" (id, username) values (1,'hugo');
	INSERT INTO theme (name, accent_500, primary_500, warn_500) VALUES ('standard theme', '#123456', '#123456', '#123456');
	`
	if _, err := conn.Exec(ctx, sql); err != nil {
		t.Fatalf("adding example data: %v", err)
	}

	if err := <-done; err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestBigQuery(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("Postgres Test")
	}

	ctx := t.Context()

	tp, err := newTestPostgres(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	flow, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
	if err != nil {
		t.Fatalf("NewSource(): %v", err)
	}

	conn, err := tp.conn(ctx)
	if err != nil {
		t.Fatalf("create connection: %v", err)
	}

	count := 2_000
	keys := make([]dskey.Key, count)
	expected := make(map[dskey.Key][]byte)
	for i := range count {
		keys[i], _ = dskey.FromParts("user", 1, "username")
		expected[keys[i]] = []byte(`"hugo"`)

		sql := `INSERT INTO "user" (username) values ('hugo');`
		if _, err := conn.Exec(ctx, sql); err != nil {
			t.Fatalf("adding user %d: %v", i+1, err)
		}
	}

	got, err := flow.Get(ctx, keys...)
	if err != nil {
		t.Errorf("Sending request with %d fields returns: %v", count, err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got != expected: %v, %v", got[keys[0]], expected[keys[0]])
	}
}

type testPostgres struct {
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource

	Env map[string]string

	pgxConfig *pgx.ConnConfig
}

func newTestPostgres(ctx context.Context) (tp_ *testPostgres, err error) {
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

	tp := &testPostgres{
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

func (tp *testPostgres) Close() error {
	if err := tp.dockerPool.Purge(tp.dockerResource); err != nil {
		return fmt.Errorf("purge postgres container: %w", err)
	}
	return nil
}

func (tp *testPostgres) conn(ctx context.Context) (*pgx.Conn, error) {
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

func (tp *testPostgres) addSchema(ctx context.Context) error {
	conn, err := tp.conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}
	defer conn.Close(ctx)

	schema, err := os.ReadFile("../meta/dev/sql/schema_relational.sql")
	if err != nil {
		return fmt.Errorf("reading schema file: %w", err)
	}

	if _, err := conn.Exec(ctx, string(schema)); err != nil {
		return fmt.Errorf("adding schema: %w", err)
	}
	return nil
}

func (tp *testPostgres) cleanup(t *testing.T) {
	t.Cleanup(func() {
		ctx := context.Background()
		if err := tp.reset(ctx); err != nil {
			t.Logf("Cleanup database: %v", err)
		}
	})
}

func (tp *testPostgres) reset(ctx context.Context) error {
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

func (tp *testPostgres) addExampleData(ctx context.Context) error {
	exampleData, err := os.ReadFile("../meta/dev/sql/example_transactional.sql")
	if err != nil {
		return fmt.Errorf("reading example file: %w", err)
	}

	conn, err := tp.conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}

	if _, err := conn.Exec(ctx, string(exampleData)); err != nil {
		return fmt.Errorf("adding example data: %w", err)
	}

	return nil
}
