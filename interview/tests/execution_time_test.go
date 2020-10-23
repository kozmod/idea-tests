package tests

import (
	"testing"
	"time"
)

func TestChanAssigment(t *testing.T) {
	worker := func() <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			time.Sleep(3 * time.Second)
			ch <- struct{}{}
		}()
		return ch
	}

	timeStart := time.Now()

	_, _ = <-worker(), <-worker()

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
}

func TestChanAssigment_fix(t *testing.T) {
	ch := make(chan struct{})
	worker := func() {
		go func() {
			time.Sleep(3 * time.Second)
			ch <- struct{}{}
		}()
	}

	timeStart := time.Now()

	worker()
	worker()

	_, _ = <-ch, <-ch

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
}
