package std

import (
	"context"
	"github.com/pkg/errors"
	"testing"
	"time"
)

// TestCtxDeadline 测试context超时处理
func TestCtxDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	f := func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				//return ctx.Err()
				return errors.Wrapf(ctx.Err(), "超时错误！")
			default:
			}
		}
	}
	err := f(ctx)
	if errors.As(err, &context.DeadlineExceeded) {
		t.Log("get DeadlineExceeded err")
	} else if err != nil {
		t.Fatal(err)
	}
}

// TestContextCycleWithTimeout 测试WithTimeout会不会覆盖之前设置的超时时间
func TestContextCycleWithTimeout(t *testing.T) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	time.Sleep(1*time.Second)
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	time.Sleep(2*time.Second)
	select {
	case <- ctx.Done():
		t.Log("WithTimeout不会覆盖之前设置的超时时间")
	default:
		t.Log("WithTimeout会覆盖之前设置的超时时间")
	}
}

// TestContextWithDeadline WithDeadline会不会覆盖之前设置的超时时间
func TestContextWithDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	time.Sleep(1*time.Second)
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	time.Sleep(2*time.Second)
	select {
	case <- ctx.Done():
		t.Log("WithDeadline不会覆盖之前设置的超时时间")
	default:
		t.Log("WithDeadline会覆盖之前设置的超时时间")
	}
}
