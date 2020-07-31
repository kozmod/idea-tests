package function

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

const (
	a = "a"
	b = "b"
)

type someFn func() string

type fnStruct struct {
	fn someFn
}

func TestFunctionAsArg_Mutation(t *testing.T) {
	s := fnStruct{fn: func() string {
		return a
	}}
	assert.Equal(t, a, s.fn())
	s.fn = func() string {
		return b
	}
	assert.Equal(t, b, s.fn())
}
