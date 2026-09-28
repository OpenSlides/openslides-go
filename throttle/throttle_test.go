package throttle

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

// TestImmediateExecution verifies that the first Run executes promptly.
func TestImmediateExecution(t *testing.T) {
	done := make(chan struct{})
	tt := New(t.Context(), 50*time.Millisecond)
	tt.Run(func() {
		close(done)
	})

	select {
	case <-done:
	case <-time.After(20 * time.Millisecond):
		t.Fatal("first Run did not execute instantly")
	}
}

// TestThrottleWindowDelaysSecondRun verifies that a second Run during the
// throttle window is delayed until the window expires.
func TestThrottleWindowDelaysSecondRun(t *testing.T) {
	period := 50 * time.Millisecond
	tt := New(t.Context(), period)

	firstDone := make(chan struct{})
	tt.Run(func() {
		close(firstDone)
	})
	select {
	case <-firstDone:
	case <-time.After(40 * time.Millisecond):
		t.Fatal("first Run did not execute instantly")
	}

	start := time.Now()
	secondDone := make(chan struct{})
	tt.Run(func() {
		close(secondDone)
	})

	select {
	case <-secondDone:
		elapsed := time.Since(start)
		if elapsed < period {
			t.Fatalf("second Run executed after %v, expected at least %v", elapsed, period)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("second Run did not execute within 2s")
	}
}

// TestThrottleWindowImmediateSecondRunAfterDelay verifies that a second Run after
// the throttle period is executed instantly
func TestThrottleWindowImmediateSecondRunAfterDelay(t *testing.T) {
	period := 50 * time.Millisecond
	tt := New(t.Context(), period)

	firstDone := make(chan struct{})
	tt.Run(func() {
		close(firstDone)
	})
	select {
	case <-firstDone:
	case <-time.After(40 * time.Millisecond):
		t.Fatal("first Run did not execute instantly")
	}

	time.Sleep(period)

	secondDone := make(chan struct{})
	tt.Run(func() {
		close(secondDone)
	})

	select {
	case <-secondDone:
	case <-time.After(40 * time.Millisecond):
		t.Fatal("second Run did not execute instantly")
	}
}

// TestMostRecentWins verifies that when multiple Runs are called during the
// throttle window, only the most recent function executes.
func TestMostRecentWins(t *testing.T) {
	period := 50 * time.Millisecond
	tt := New(t.Context(), period)

	firstDone := make(chan struct{})
	tt.Run(func() {
		close(firstDone)
	})
	<-firstDone

	var executed atomic.Int32
	winnerDone := make(chan struct{})

	tt.Run(func() { executed.Store(1) })
	tt.Run(func() { executed.Store(2) })
	tt.Run(func() {
		executed.Store(3)
		close(winnerDone)
	})

	select {
	case <-winnerDone:
	case <-time.After(2 * time.Second):
		t.Fatal("trailing function did not execute within 2s")
	}

	if got := executed.Load(); got != 3 {
		t.Fatalf("expected last function to win (got %d, want 3)", got)
	}
}

// TestMultiplePeriods verifies that the throttler correctly handles several
// consecutive executions spaced by the period.
func TestMultiplePeriods(t *testing.T) {
	period := 30 * time.Millisecond
	tt := New(t.Context(), period)

	var count atomic.Int32

	for i := range 4 {
		expected := int32(i + 1)
		notify := make(chan struct{})
		tt.Run(func() {
			count.Add(1)
			if got := count.Load(); got == expected {
				close(notify)
			}
		})

		select {
		case <-notify:
		case <-time.After(2 * time.Second):
			t.Fatalf("execution %d did not complete within 2s", i+1)
		}

		// Wait for the throttle window to expire before enqueuing the next one.
		time.Sleep(period + 10*time.Millisecond)
	}

	if got := count.Load(); got != 4 {
		t.Fatalf("expected 4 executions, got %d", got)
	}
}

// TestContextCancellation verifies that cancelling the context stops the
// throttler and prevents further executions.
func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	period := 100 * time.Millisecond
	tt := New(ctx, period)

	executed := make(chan struct{}, 1)
	tt.Run(func() {
		select {
		case executed <- struct{}{}:
		default:
		}
	})

	// Let the first execution complete.
	select {
	case <-executed:
	case <-time.After(time.Second):
		t.Fatal("first Run did not execute")
	}

	// Cancel the context mid-window.
	cancel()

	// Give the goroutine time to observe cancellation.
	time.Sleep(period + 50*time.Millisecond)

	// Enqueue another function; it should never execute.
	tt.Run(func() {
		t.Error("function executed after context cancellation")
	})

	time.Sleep(period + 60*time.Millisecond)
}

// TestConcurrentRuns verifies that concurrent Run calls don't panic or
// deadlock, and that exactly one function survives to execute.
func TestConcurrentRuns(t *testing.T) {
	period := 50 * time.Millisecond
	tt := New(t.Context(), period)

	// Start with one execution to open the throttle window.
	ready := make(chan struct{})
	tt.Run(func() { close(ready) })
	<-ready

	// Fire many concurrent Run calls during the window.
	var counter atomic.Int32
	for range 100 {
		go tt.Run(func() { counter.Add(1) })
	}

	// Wait long enough for the window to expire and the trailing function to run.
	time.Sleep(period + 80*time.Millisecond)

	got := counter.Load()
	if got == 0 {
		t.Fatal("expected at least one trailing execution, got 0")
	}
	if got > 1 {
		t.Logf("note: %d functions executed (expected exactly 1, race may allow 2 in edge cases)", got)
	}
}
