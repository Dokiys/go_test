package std

import (
	"context"
	"runtime"
	"testing"
	"time"
)

// TestGoroutineTickerBreak	select-case无法已经开始运行的goroutine
func TestGoroutineTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {}
		done <- true
	}()

	select {
	case <-done:
		t.Log("go func Done")
	case  <-ticker.C:
		t.Log("Tick at")
	}
	t.Logf("running goroutine num: %d",runtime.NumGoroutine())
	select{}
}

// TestGoroutineCtxBreak ctx的方式也一样无法停止已经运行的goroutine
func TestGoroutineCtxBreak(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {}

		select {
		case <-ctx.Done():
			return
		}
	}(ctx)

	t.Logf("running goroutine num: %d",runtime.NumGoroutine())
	select {}
}