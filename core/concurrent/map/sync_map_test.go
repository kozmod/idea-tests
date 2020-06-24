package _map

import (
	"sync"
	"testing"

	"github.com/kozmod/idea-tests/utils/tsync"
	"github.com/stretchr/testify/assert"
)

const keySyncMap = "counter"

var mx sync.RWMutex

func TestSyncMap(t *testing.T) {
	iteration := 100_000
	assert.NotPanics(t, func() {
		var syncMap sync.Map
		syncMap.Store(keySyncMap, 0)

		tsync.MultiSubmit(iteration, func() {
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
