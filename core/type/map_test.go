package _type

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapLen(t *testing.T) {
	m := make(map[string]uint)
	assert.Equal(t, len(m), 0)

	m["0"] = 0
	m["1"] = 0
	m["2"] = 0
	assert.Equal(t, len(m), 3)

	m = make(map[string]uint, 10)
	assert.Equal(t, len(m), 0)

	m["0"] = 0
	m["1"] = 0
	m["2"] = 0
	assert.Equal(t, len(m), 3)
}
