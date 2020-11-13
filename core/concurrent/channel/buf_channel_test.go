package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

const year = 8640 * time.Hour

var (
	first  = time.Now()
	second = first.Add(1 * year)
	third  = first.Add(2 * year)
)

func TestBuffChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	bufferedChan := make(chan time.Time, 3)
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
