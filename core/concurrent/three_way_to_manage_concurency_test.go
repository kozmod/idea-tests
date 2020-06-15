package concurrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2) //task count wait to do
	go func() {
		wg.Done() // finish task1
		fmt.Println("goroutine 1 done ")
	}()
	go func() {
		wg.Done() // finish task2
		fmt.Println("goroutine 2 done ")
	}()
	wg.Wait() // wait for tasks to be don
}

func TestChan(t *testing.T) {
	exit := make(chan bool)
	go func() {
		for {
			select {
			case <-exit:
				fmt.Println("Exit")
				return
			default:
				fmt.Println("Monitoring")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Notify Exit")
	exit <- true
	//keep main goroutine alive
	time.Sleep(5 * time.Second)
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go A(ctx, "Call One") //mock clients calls A
	time.Sleep(3 * time.Second)
	fmt.Println("client release connection, need to notify A, B exit")
	cancel() //mock client exit, and pass the signal, ctx.Done() gets the signal
	time.Sleep(3 * time.Second)
}

func A(ctx context.Context, name string) {
	go B(ctx, name) // A calls B
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A Exit")
			return
		default:
			fmt.Println(name, "A do something")
			time.Sleep(2 * time.Second)
		}
	}
}
func B(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B Exit")
			return
		default:
			fmt.Println(name, "B do something")
			time.Sleep(2 * time.Second)
		}
	}
}
