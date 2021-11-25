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
