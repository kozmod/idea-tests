package channel

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"sync"
	"testing"
)

const (
	first  = "first"
	second = "second"
	third  = "third"
)

func TestBuffChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	bufferedChan := make(chan string, 3)
	go func() {
		fmt.Println("Sending..")
		bufferedChan <- first
		bufferedChan <- second
		bufferedChan <- third
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		fmt.Println("Receiving..")
		assert.Equal(t, first, <-bufferedChan)
		assert.Equal(t, second, <-bufferedChan)
		assert.Equal(t, third, <-bufferedChan)
		wg.Done()
	}()
	wg.Wait()
}
