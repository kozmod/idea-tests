package tests

import (
	"fmt"
	"sync"
	"testing"
)

type Counter struct {
	value int
}

func TestRangePointer(t *testing.T) {
	var res = make([]*Counter, 3)
	for i, a := range []Counter{{1}, {2}, {3}} {
		res[i] = &a
	}
	fmt.Println("res:", res[0].value, res[1].value, res[2].value)
}

func TestRangeVal(t *testing.T) {
	wg := sync.WaitGroup{}
	data := []string{"one", "two", "three"}
	for _, v := range data {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
