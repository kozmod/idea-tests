package tests

import (
	"fmt"
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
	fmt.Println(counter)
}
