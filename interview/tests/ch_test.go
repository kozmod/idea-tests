package tests

import (
	"fmt"
	"sync"
	"testing"
)

func TestCh_1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan string)
	close(ch)
	go func() {
		val, ok := <-ch
		fmt.Println(val, ok)
		wg.Done()
	}()
	wg.Wait()
}

func TestCh_2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan string)
	close(ch)
	go func() {
		ch <- "val"
		wg.Done()
	}()
	wg.Wait()
	val, ok := <-ch
	fmt.Println(val, ok)
}

func TestCh_3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	var ch chan int
	go func() {
		ch <- 123
		wg.Done()
	}()
	wg.Wait()
}

func TestCh_4(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	var ch chan int
	go func() {
		val, ok := <-ch
		fmt.Println(val, ok)
		wg.Done()
	}()
	wg.Wait()
}

func TestCh_5(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	var ch chan int
	go func() {
		close(ch)
		wg.Done()
	}()
	wg.Wait()
}
