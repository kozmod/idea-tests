package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	s := fnStruct{}
	assert.Nil(t, s.fn)

	s.fn = func() string {
		return a
	}
	assert.Equal(t, a, s.fn())

	s.fn = func() string {
		return b
	}
	assert.Equal(t, b, s.fn())
}
