package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/ory/dockertest/v4"
)

type testRedis struct {
	addr string
	Env  map[string]string
}

func newTestRedis(t *testing.T) *testRedis {
	t.Helper()

	if testing.Short() {
		t.Skip("Redis Test")
	}

	pool := dockertest.NewPoolT(t, "")

	redis := pool.RunT(t, "redis",
		dockertest.WithTag("6.2"),
	)

	port := redis.GetPort("6379/tcp")
	addr := ":" + port

	err := pool.Retry(t.Context(), 30*time.Second, func() error {
		conn, err := redigo.DialContext(t.Context(), "tcp", addr)
		if err != nil {
			return fmt.Errorf("create connection to redis: %w", err)
		}
		defer conn.Close()

		if _, err := conn.Do("PING"); err != nil {
			return fmt.Errorf("send ping: %w", err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("redis not ready: %v", err)
	}

	tr := &testRedis{
		addr: addr,
		Env: map[string]string{
			"MESSAGE_BUS_PORT": port,
		},
	}

	return tr
}

func (tr *testRedis) conn(ctx context.Context) (redigo.Conn, error) {
	conn, err := redigo.DialContext(ctx, "tcp", tr.addr)
	if err != nil {
		return nil, fmt.Errorf("creating test connection: %w", err)
	}

	return conn, nil
}
