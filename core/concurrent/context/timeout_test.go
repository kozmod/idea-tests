package context

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kozmod/idea-tests/utils/tsync"
)

const (
	first  = "first"
	second = "second"
	third  = "third"
)

func TestDoneByTimeout(t *testing.T) {
	m := tsync.OrderExecuteAll(
		func() interface{} {
			return ex(1, 4, first)
		},
		func() interface{} {
			return ex(4, 3, second)
		},
		func() interface{} {
			return ex(4, 4, third)
		},
	)
	fmt.Println("  <------------------------------>  ")
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%d - %s", k, v))
	}

	fmt.Println("  <------------------------------>  ")
}

func ex(extime time.Duration, timeout time.Duration, val string) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	fmt.Println(val + " server: ex started")
	defer fmt.Println(val + " server: ex ended")

	select {
	case <-time.After(extime * time.Second):
		fmt.Println(val + " - hello")
		return fmt.Sprintf("%s - done  \n {extime %d, tomeout %d}", val, extime, timeout)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(val+" server err:", err)
		return fmt.Sprintf("%s - cenceled \n {extime %d, tomeout %d}", val, extime, timeout)
	}
}
