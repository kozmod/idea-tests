package concurrent

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	var (
		x   = 0
		req = 10_000
	)
	var wg sync.WaitGroup
	wg.Add(req)
	for i := 0; i < req; i++ {
		go func() {
			x++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
	assert.NotEqual(t, req, x)
}

func TestNoRaceCondition(t *testing.T) {
	var (
		x   = 0
		req = 10_000
	)
	var memoryAccess sync.Mutex
	var wg sync.WaitGroup
	wg.Add(req)
	for i := 0; i < req; i++ {
		go func() {
			memoryAccess.Lock()
			x++
			memoryAccess.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
	assert.Equal(t, req, x)
}
