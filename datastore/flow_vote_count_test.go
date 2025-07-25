package datastore_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/environment"
)

func TestVoteCountSourceGet(t *testing.T) {
	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
		w.(http.Flusher).Flush()
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	ctx := t.Context()

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	flow := datastore.NewFlowVoteCount(env)
	eventer := func() (<-chan time.Time, func() bool) { return make(chan time.Time), func() bool { return true } }

	waitForResponse(ctx, flow, func() {
		go flow.Connect(ctx, eventer, func(error) {})
	})

	key1 := dskey.MustKey("poll/1/live_votes")

	t.Run("no data from vote-service", func(t *testing.T) {
		got, err := flow.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() without any data returned %s, expected nil", got[key1])
		}
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, flow, func() {
			sender <- `{"1":{"42":null}}`
		})

		got, err := flow.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != `{"42":null}` {
			t.Errorf(`Get() after first data returned %s, expected {"42":null}`, got[key1])
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, flow, func() {
			sender <- `{"1":{"43":null}}`
		})

		got, err := flow.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != `{"42":null,"43":null}` {
			t.Errorf("Get() after first data returned `%s`, expected `43`", got[key1])
		}
	})

	t.Run("again data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, flow, func() {
			sender <- `{"1":{"44":null}}`
		})

		got, err := flow.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != `{"42":null,"43":null,"44":null}` {
			t.Errorf("Get() after first data returned `%s`, expected `44`", got[key1])
		}
	})

	t.Run("receive null", func(t *testing.T) {
		waitForResponse(ctx, flow, func() {
			sender <- `{"1":null}`
		})

		got, err := flow.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() after receiving 0 returned %s, expected nil", got[key1])
		}
	})
}

func TestVoteCountSourceUpdate(t *testing.T) {
	ctx := t.Context()

	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
		w.(http.Flusher).Flush()

		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	flow := datastore.NewFlowVoteCount(env)
	eventer := func() (<-chan time.Time, func() bool) { return make(chan time.Time), func() bool { return true } }

	waitForResponse(ctx, flow, func() {
		go flow.Connect(ctx, eventer, func(error) {})
	})

	key1 := dskey.MustKey("poll/1/live_votes")

	t.Run("no data from vote-service", func(t *testing.T) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond)
		defer cancel()

		flow.Update(ctxTimeout, func(map[dskey.Key][]byte, error) {
			t.Fatalf("Update was called")
		})
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		got, err := updateResult(ctx, flow, func() {
			sender <- `{"1":{"42":null}}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: []byte(`{"42":null}`)}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		got, err := updateResult(ctx, flow, func() {
			sender <- `{"1":{"43":null}}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: []byte(`{"42":null,"43":null}`)}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("receive null", func(t *testing.T) {
		got, err := updateResult(ctx, flow, func() {
			sender <- `{"1":null}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: nil}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})
}

func TestReconnect(t *testing.T) {
	msg := `{"1":[23]}`
	sender := make(chan struct{})
	var counter int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
		w.(http.Flusher).Flush()
		counter++
		<-sender
	}))

	ctx := t.Context()

	event := make(chan time.Time)
	close(event)
	eventer := func() (<-chan time.Time, func() bool) {
		return event, func() bool { return false }
	}

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	flow := datastore.NewFlowVoteCount(env)
	go flow.Connect(ctx, eventer, func(error) {})

	sender <- struct{}{} // Close connection so there is a reconnect
	sender <- struct{}{} // Close connection again

	if counter < 2 {
		t.Errorf("Got %d connections, expected 2", counter)
	}
}

func TestGetWithoutConnect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := datastore.NewFlowVoteCount(env)

	key := dskey.MustKey("poll/1/live_votes")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	_, err := source.Get(ctxTimeout, key)
	if err != context.DeadlineExceeded {
		t.Fatalf("Update: %v, expected context.DeadlineExceeded", err)
	}
}

// waitForResponse calls the given function and waits until the data is
// processed.
func waitForResponse(ctx context.Context, flow *datastore.FlowVoteCount, fn func()) {
	updateResult(ctx, flow, fn)
}

// updateResult returns the return values from flow.Update after the given function is processed.
func updateResult(ctx context.Context, flow *datastore.FlowVoteCount, fn func()) (map[dskey.Key][]byte, error) {
	type dataErr struct {
		data map[dskey.Key][]byte
		err  error
	}

	myCtx, cancel := context.WithCancel(ctx)

	got := dataErr{}
	go flow.Update(myCtx, func(v map[dskey.Key][]byte, err error) {
		got = dataErr{v, err}
		cancel()
	})

	fn()

	<-myCtx.Done()

	return got.data, got.err
}

func parseURL(raw string) (host, port, protocol string) {
	parsed, err := url.Parse(raw)
	if err != nil {
		panic(fmt.Sprintf("parsing url %s: %v", raw, err))
	}

	return parsed.Hostname(), parsed.Port(), parsed.Scheme
}
