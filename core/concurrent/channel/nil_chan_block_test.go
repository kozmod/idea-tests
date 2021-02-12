package channel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFor_NilChannelBlocks(t *testing.T) {
	stop := make(chan bool)

	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//// do other work
	go func() {
		time.Sleep(5 * time.Second)
		stop <- true
	}()

	select {
	case <-ch:
		t.Fail()
		break
	case st := <-stop:
		assert.True(t, st)
		break
	}
}

func Test_NilChannelBlocks_Solve_WithSignalChannel(t *testing.T) {
	doWork := func(done <-chan interface{}, strings <-chan string,
	) <-chan interface{} { // 1
		terminared := make(chan interface{})

		go func() {
			defer close(terminared)

			for {
				select {
				case <-done: // 2
					return
				case s := <-strings:
					fmt.Println(s)
				}
			}
		}()
		return terminared
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { // 3
		// Cancel the operation after 1 second
		time.Sleep(1 * time.Second)
		close(done)
	}()

	<-terminated //4
	fmt.Println("Done.")
}
