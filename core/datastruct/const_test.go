package datastruct

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type X int

const (
	a X = iota
	b
	c
)

func (d X) String() string {
	return [...]string{"a", "b", "c"}[d]
}

func Test_1(t *testing.T) {
	assert.Equal(t, c.String(), "c")
}
