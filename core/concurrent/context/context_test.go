package context

import (
	"context"
	"fmt"
	"github.com/kozmod/idea-tests/core/utils/tsync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	m := tsync.OrderExecuteAll(
		func() interface{} {
			return do(1, 4, "first")
		},
		func() interface{} {
			return do(4, 3, "second")
		},
		func() interface{} {
			return do(4, 4, "third")
		},
	)
	fmt.Println("  <------------------------------>  ")
	fmt.Println(m)
	fmt.Println("  <------------------------------>  ")

}

func do(extime time.Duration, timeout time.Duration, val string) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	fmt.Println(val + " server: do started")
	defer fmt.Println(val + " server: do ended")

	select {
	case <-time.After(extime * time.Second):
		fmt.Println(val + " - hello")
		return val + " - done"
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(val+" server err:", err)
		return val + " - canceled"
	}
}
