package interesting_usage

import (
	"fmt"
	"testing"
	"time"
)

func stopChanWorker(die chan bool, i int) {
	for {
		work := make(chan bool)

		go func(x chan bool) {
			fmt.Printf("Start %d \n", i)
			time.Sleep(1 * time.Second)
			x <- true
		}(work)

		select {
		case <-work:
			fmt.Printf("Stop %d \n", i)
		case <-die:
			return
		}
	}
}

func TestStopAll(t *testing.T) {
	die := make(chan bool)

	for i := 0; i < 5; i++ {
		go stopChanWorker(die, i)
	}
	time.Sleep(1 * time.Second)
	// Остановить всех stopChanWorker'ов.
	close(die)
}
