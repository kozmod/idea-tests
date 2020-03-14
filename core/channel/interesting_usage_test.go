package channel

import (
	"fmt"
	"testing"
	"time"
)

func worker(start chan bool, num int) {
	fmt.Println(num)
	<-start
	fmt.Printf("Start %d \n", num)
}

//В этом примере сотня горутин запускается, ждет передачи данных через канал start (или его закрытия).
//В случае, когда он будет закрыт, все горутины запустятся.
func TestWorker(t *testing.T) {
	start := make(chan bool)

	for i := 0; i < 10; i++ {
		go worker(start, i)
	}
	time.Sleep(1 * time.Second)
	close(start)
	time.Sleep(1 * time.Second)
}
