package context

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func TestTicker2(t *testing.T) {
	done := make(chan bool)
	done2 := make(chan bool)

	go func() {
		//for {
		select {
		case <-done:
			fmt.Println("d")
		case <-done2:
			fmt.Println("d2")
			break
		}
		fmt.Println("end")
		//}
		fmt.Println("END")
	}()

	time.Sleep(2 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
	time.Sleep(2 * time.Second)
	done2 <- true
	time.Sleep(2 * time.Second)
}
