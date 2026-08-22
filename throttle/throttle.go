package throttle

import (
	"context"
	"sync"
	"time"
)

// Throttler limits how often enqueued functions are executed.
type Throttler struct {
	nextFn func()
	cond   *sync.Cond
}

// New returns a new Throttler
func New(ctx context.Context, period time.Duration) *Throttler {
	t := &Throttler{
		cond:   sync.NewCond(&sync.Mutex{}),
		nextFn: nil,
	}

	go t.loop(ctx, period)
	go (func() {
		<-ctx.Done()
		t.cond.Signal()
	})()

	return t
}

// Run enqueues fn for throttled execution. If a function is already pending
// (waiting for the throttle window to expire), it is replaced by fn so that
// only the most recently enqueued function runs.
func (tt *Throttler) Run(fn func()) {
	tt.cond.L.Lock()
	tt.nextFn = fn
	tt.cond.L.Unlock()
	tt.cond.Signal()
}

func (tt *Throttler) loop(ctx context.Context, period time.Duration) {
	for {
		tt.cond.L.Lock()
		for tt.nextFn == nil {
			tt.cond.Wait()
			if ctx.Err() != nil {
				return
			}
		}

		tt.nextFn()
		tt.nextFn = nil

		tt.cond.L.Unlock()
		time.Sleep(period)
	}
}
