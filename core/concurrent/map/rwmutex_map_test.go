package _map

import (
	"github.com/kozmod/idea-tests/utils/tsync"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const keyRwMutex = "counter"

type rwmutexMap struct {
	mx sync.RWMutex
	m  map[string]int
}

func NewRWMutexMap() *rwmutexMap {
	return &rwmutexMap{
		m: make(map[string]int),
	}
}

func (c *rwmutexMap) Load(key string) (int, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *rwmutexMap) Store(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *rwmutexMap) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key]++
}

func TestRWMutexMap(t *testing.T) {
	iteration := 100_000
	assert.NotPanics(t, func() {
		threadSafeMap := NewRWMutexMap()
		threadSafeMap.Store(keyRwMutex, 0)

		tsync.MultiSubmit(iteration, func() {
			threadSafeMap.Inc(keyRwMutex)
		})

		res, _ := threadSafeMap.Load(keyRwMutex)
		assert.Equal(t, iteration, res)
	}, "The code did panic")
}
