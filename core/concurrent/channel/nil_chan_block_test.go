package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestForRange_NilChannelBlocks(t *testing.T) {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	//// do other work

	go func() {
		time.Sleep(5 * time.Second)
		panic("c")
	}()

	// get first result - BLOCK
	fmt.Println("result:", <-ch)
}
