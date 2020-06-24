package tsync_test

import (
	"testing"

	. "github.com/kozmod/idea-tests/utils/tsync"
	"github.com/stretchr/testify/assert"
)

const a = "a"
const one = 1
const two = 2
const three = 3

func TestKeyExecuteAll(t *testing.T) {
	m := KeyExecuteAll(
		map[interface{}]func() interface{}{
			one: func() interface{} {
				return one
			},
			a: func() interface{} {
				return a
			},
			three: func() interface{} {
				return three
			},
			two: func() interface{} {
				return two
			},
		})
	assert.Equal(t, a, m[a])
	assert.Equal(t, one, m[one])
	assert.Equal(t, two, m[two])
	assert.Equal(t, three, m[three])
}

func TestOrderExecuteAll(t *testing.T) {
	m := OrderExecuteAll(
		func() interface{} {
			return a
		},
		func() interface{} {
			return one
		},
		func() interface{} {
			return two
		},
	)
	assert.Equal(t, a, m[0])
	assert.Equal(t, one, m[1])
	assert.Equal(t, two, m[2])
}
