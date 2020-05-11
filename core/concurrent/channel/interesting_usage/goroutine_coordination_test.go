package interesting_usage

import (
	"fmt"
	"testing"
	"time"
)

//В этом примере сотня горутин запускается, ждет передачи данных через канал start (или его закрытия).
//В случае, когда он будет закрыт, все горутины запустятся.
func TestStartCoordination(t *testing.T) {
	start := make(chan bool)

	for i := 0; i < 10; i++ {
		go waitWorker(start, i)
	}
	time.Sleep(1 * time.Second)
	close(start)
	time.Sleep(1 * time.Second)
}

func waitWorker(start chan bool, num int) {
	fmt.Printf("Start %d \n", num)
	<-start
	fmt.Printf("Stop %d \n", num)
}
