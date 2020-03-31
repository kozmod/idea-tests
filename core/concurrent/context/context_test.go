package context

import (
	"context"
	"fmt"
	"github.com/kozmod/idea-tests/core/utils/tsync"
	"testing"
	"time"
)

const (
	first  = "first"
	second = "second"
	third  = "third"
)

func TestContext(t *testing.T) {
	m := tsync.OrderExecuteAll(
		func() interface{} {
			return do(1, 4, first)
		},
		func() interface{} {
			return do(4, 3, second)
		},
		func() interface{} {
			return do(4, 4, third)
		},
	)
	fmt.Println("  <------------------------------>  ")
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%d - %s", k, v))
	}

	fmt.Println("  <------------------------------>  ")
}

func do(extime time.Duration, timeout time.Duration, val string) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	fmt.Println(val + " server: do started")
	defer fmt.Println(val + " server: do ended")

	select {
	case <-doSth(extime * time.Second):
		fmt.Println(val + " - hello")
		return fmt.Sprintf("%s - done  \n {extime %d, tomeout %d}", val, extime, timeout)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(val+" server err:", err)
		return fmt.Sprintf("%s - cenceled \n {extime %d, tomeout %d}", val, extime, timeout)
	}
}

func doSth(duration time.Duration) <-chan time.Time {
	return time.After(duration)
}
