package throttle

import (
	"context"
	"time"
)

// Throttler limits how often enqueued functions are executed.
type Throttler struct {
	in chan func()
}

// New returns a new Throttler
func New(ctx context.Context, period time.Duration) *Throttler {
	t := &Throttler{
		in: make(chan func(), 1),
	}

	go t.loop(ctx, period)
	return t
}

// Run enqueues fn for throttled execution. If a function is already pending
// (waiting for the throttle window to expire), it is replaced by fn so that
// only the most recently enqueued function runs.
func (tt *Throttler) Run(fn func()) {
	for {
		select {
		case tt.in <- fn:
			return
		default:
		}

		select {
		case <-tt.in:
		default:
		}
	}
}

func (tt *Throttler) loop(ctx context.Context, period time.Duration) {
	var timer *time.Timer
	var timerC <-chan time.Time
	var nextFn func()

	for {
		select {
		case <-ctx.Done():
			if timer != nil {
				timer.Stop()
			}
			return

		case fn := <-tt.in:
			if timer == nil {
				timer = time.NewTimer(0)
			} else {
				timer.Reset(period)
			}

			nextFn = fn
			if timerC == nil {
				timerC = timer.C
			}

		case <-timerC:
			nextFn()
			nextFn = nil
			timerC = nil
		}
	}
}
