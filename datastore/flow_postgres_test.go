package datastore_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/pgtest"
	"github.com/OpenSlides/openslides-go/environment"
)

func TestFlowPostgres(t *testing.T) {
	ctx := t.Context()

	t.Parallel()
	if testing.Short() {
		t.Skip("Postgres Test")
	}

	tp, err := pgtest.NewPostgresTest(ctx)
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
			``,
			map[string][]byte{
				"theme/1/name": []byte(`"standard theme"`),
			},
		},
		{
			"different fqid",
			`
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
			`UPDATE organization SET enable_electronic_voting=true WHERE id = 1`,
			map[string][]byte{
				"organization/1/enable_electronic_voting": []byte(`true`),
			},
		},
		{
			"Timestamp",
			`UPDATE "user" set last_login='1999-01-08' WHERE id=1;`,
			map[string][]byte{
				"user/1/last_login": []byte(`915753600`),
			},
		},
		{
			"Float",
			`INSERT INTO "projector_countdown" (title, countdown_time, meeting_id) VALUES ('test countdown', 7.5, 1);`,
			map[string][]byte{
				"projector_countdown/1/countdown_time": []byte(`7.5`),
			},
		},
		{
			"String list",
			`
			INSERT INTO history_entry (entries) VALUES ('{"entry1","entry2"}');
			`,
			map[string][]byte{
				"history_entry/1/entries": []byte(`["entry1","entry2"]`),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			tp.Cleanup(t)

			conn, err := tp.Conn(ctx)
			if err != nil {
				t.Fatalf("create connection: %v", err)
			}
			defer conn.Close(ctx)

			if _, err := conn.Exec(ctx, tt.insert); err != nil {
				t.Fatalf("adding example data: %v", pgtest.PrityPostgresError(err, tt.insert))
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

	tp, err := pgtest.NewPostgresTest(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	conn, err := tp.Conn(ctx)
	if err != nil {
		t.Fatalf("create connection: %v", err)
	}

	flow, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
	if err != nil {
		t.Fatalf("NewFlowPostgres(): %v", err)
	}

	keys := []dskey.Key{
		dskey.MustKey("user/300/username"),
		dskey.MustKey("theme/300/name"),
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
	INSERT INTO "user" (id, username) VALUES (300,'hugo');
	INSERT INTO theme (id, name) VALUES (300,'standard theme');
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

	tp, err := pgtest.NewPostgresTest(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	flow, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env))
	if err != nil {
		t.Fatalf("NewSource(): %v", err)
	}

	conn, err := tp.Conn(ctx)
	if err != nil {
		t.Fatalf("create connection: %v", err)
	}

	count := 2_000
	keys := make([]dskey.Key, count)
	expected := make(map[dskey.Key][]byte)
	for i := range count {
		keys[i], _ = dskey.FromParts("user", i+2, "username")
		expected[keys[i]] = []byte(`"hugo"`)

		sql := `INSERT INTO "user" (username) values ('hugo');`
		if _, err := conn.Exec(ctx, sql); err != nil {
			t.Fatalf("adding user %d: %v", i+2, err)
		}
	}

	got, err := flow.Get(ctx, keys...)
	if err != nil {
		t.Errorf("Sending request with %d fields returns: %v", count, err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got != expected: %s, %s", got[keys[0]], expected[keys[0]])
	}
}
