package _map

import (
	"github.com/kozmod/idea-tests/core/utils/tsync"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const keyMutex = "counter"

type mutexMap struct {
	mx sync.Mutex
	m  map[string]int
}

func NewMutexMap() *mutexMap {
	return &mutexMap{
		m: make(map[string]int),
	}
}

func (c *mutexMap) Load(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *mutexMap) Store(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *mutexMap) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

func TestMutexMap(t *testing.T) {
	iteration := 100_000
	assert.NotPanics(t, func() {
		threadSafeMap := NewMutexMap()
		threadSafeMap.Store(keyMutex, 0)

		tsync.MultiSubmit(iteration, func() {
			threadSafeMap.Inc(keyMutex)
		})

		res, _ := threadSafeMap.Load(keyMutex)
		assert.Equal(t, iteration, res)
	}, "The code did panic")
}

func TestMap(t *testing.T) {
	assert.Panics(t, func() {
		m := make(map[string]int)
		m[keyMutex] = 0
		for i := 0; i < 100_00; i++ {
			go func() {
				m[keyMutex]++
			}()
		}
	}, "The code did not panic")
}
