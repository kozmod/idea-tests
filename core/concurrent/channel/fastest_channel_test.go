package channel

import (
	"fmt"
	"testing"
)

func TestGetFirsResult(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			go func(c chan<- int, i int) {
				res := Add(i)
				c <- res
			}(ch, i)
		}
	}()
	res := <-ch //blocking, before get first result
	//close(ch) - writing to close channel produce a panic
	fmt.Println(res)
}

func Add(num int) int {
	return num + 5
}
