package channel

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestForRange_3(t *testing.T) {
	consumer := make(chan int)
	go func(c chan<- int) {
		i := 0
		for {
			consumer <- i
			i++
			if i == 10 {
				break
			}
		}
		close(c)
	}(consumer)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			j, more := <-consumer
			if more {
				fmt.Println("received", j)
			} else {
				fmt.Println("received all val")
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}

func TestForRange_4(t *testing.T) {
	send := func(c chan<- int) {
		for i := 1; i < 6; i++ {
			time.Sleep(time.Second)
			c <- i
		}
		close(c)
	}

	c := make(chan int)
	go send(c)
	for value := range c {
		fmt.Println(value)
	}
}

func TestForRange_5(t *testing.T) {
	var wg sync.WaitGroup
	consumer := make(chan string, 1)
	go func() {
		consumer <- res1
		consumer <- res2
		consumer <- res3
		close(consumer)
	}()
	fmt.Println(len(consumer))
	i := 0
	for {
		res, ok := <-consumer
		i++
		if !ok {
			break
		}
		wg.Add(1)
		go func(i int) {
			log.Println(fmt.Sprintf("get res %s on iteration %d", res, i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
