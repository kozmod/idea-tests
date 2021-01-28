package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFori(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func TestFori_fix(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	var counter int32

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1000), counter)
}

func TestFori_fix_use_channel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	ch := make(chan int, 1)
	ch <- 0
	for i := 0; i < 1000; i++ {
		go func(cch chan int) {
			counter := <-cch
			counter++
			cch <- counter
			wg.Done()
		}(ch)
	}
	wg.Wait()
	assert.Equal(t, 1000, <-ch)
}

func TestFori_fix_use_mutex(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)

	var counter int
	var mux sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mux.Lock()
			counter++
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, 1000, counter)
}
