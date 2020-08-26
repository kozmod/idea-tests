package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go func() {
		defer waitGroup.Done()
		fmt.Println("1 go routine has started and waits for 1 second before broadcasting condition")
		time.Sleep(1 * time.Second)
		fmt.Println("1 go routine broadcasts condition")
		cond.Broadcast()
	}()

	go func() {
		lock.Lock()
		defer waitGroup.Done()
		fmt.Println("2 go routine has started and is waiting on condition")
		cond.Wait()
		fmt.Println("2 go routine unlocked by condition broadcast")
		lock.Unlock()
	}()

	go func() {
		cond.L.Lock()
		defer waitGroup.Done()
		fmt.Println("3 go routine has started and is waiting on condition")
		cond.Wait()
		fmt.Println("3 go routine unlocked by condition broadcast")
		cond.L.Unlock()
	}()

	fmt.Println("Main go routine starts waiting")
	waitGroup.Wait()
	fmt.Println("Main go routine ends")
}
