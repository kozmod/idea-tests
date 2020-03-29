package channel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

const (
	res1 = "exres1"
	res2 = "res2"
	res3 = "res3"
)

func TestForRange_1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	consumer := make(chan string)
	go func() {
		consumer <- res1
	}()
	go func() {
		for res := range consumer {
			// sig is a ^Config, handle it
			log.Println(fmt.Sprintf("get res: %s, and close channel", res))
			assert.Equal(t, res1, res)
			close(consumer)
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestForRange_2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	consumer := make(chan string, 1)
	go func() {
		consumer <- res1
		consumer <- res2
		consumer <- res3
	}()
	go func() {
		i := 0
		for res := range consumer {
			// sig is a ^Config, handle it
			i++
			log.Println(fmt.Sprintf("get res %s on iteration %d", res, i))
			if i == 3 {
				close(consumer)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
