package channel

import (
	"fmt"
	"testing"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func TestWithoutDeadlock_Select_default(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	default:
		fmt.Println("No goroutines available to send data", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}

func TestDeadlock_Select_without_default(t *testing.T) {
	t.Skip()
	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
