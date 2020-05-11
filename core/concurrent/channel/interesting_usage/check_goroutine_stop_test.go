package interesting_usage

import (
	"fmt"
	"testing"
	"time"
)

func stopWorker(die chan bool) {
	for {

		work := make(chan bool)

		go func(x chan bool) {
			fmt.Println("Start stopWorker")
			time.Sleep(1 * time.Second)
			x <- true
		}(work)

		select {
		// ... выполняем что-нибудь в других case
		case <-work:
			fmt.Println("Stop stopWorker itself")
			die <- true
			return
		case <-die:
			// ... выполняем необходимые действия перед завершением.
			fmt.Println("Stop stopWorker forced")
			die <- true
			return
		}
	}
}

func TestCheckStopForced(t *testing.T) {
	die := make(chan bool)
	go stopWorker(die)
	die <- true //принудительня остановка воркера

	// Ждем, пока все горутины закончат выполняться
	<-die
}

func TestCheckStop(t *testing.T) {
	die := make(chan bool)
	go stopWorker(die)

	// Ждем, пока все горутины закончат выполняться
	<-die
}
