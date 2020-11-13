package channel

import (
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
