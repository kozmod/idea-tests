package _map

import (
	"github.com/kozmod/idea-tests/core/utils/tsync"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const keySyncMap = "counter"

var mx sync.RWMutex

func TestSyncMap(t *testing.T) {
	iteration := 100_000
	assert.NotPanics(t, func() {
		var syncMap sync.Map
		syncMap.Store(keySyncMap, 0)

		tsync.IterateSubmit(iteration, func() {
			inc(keySyncMap, &syncMap)
		})

		res, ok := syncMap.Load(keySyncMap)
		assert.True(t, ok)
		assert.Equal(t, iteration, res)
	}, "The code did panic")
}

func inc(key string, syncMap *sync.Map) {
	mx.Lock()
	defer mx.Unlock()
	v, _ := syncMap.Load(key)
	val := v.(int)
	val++
	syncMap.Store(key, val)
}
