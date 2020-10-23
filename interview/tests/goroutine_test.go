package tests

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			runtime.Gosched()
			fmt.Println(s)
		}
	}

	go say("world")
	say("hello")
	time.Sleep(1)
}
