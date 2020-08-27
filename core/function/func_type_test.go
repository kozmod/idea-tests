package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StringActor func(string) (string, error)

func (act StringActor) act(s string) (string, error) {
	return act(s)
}

func newAddPrefixStringActor(prefix string) StringActor {
	return func(s string) (string, error) {
		return prefix + s, nil
	}
}

func TestFunctionAsType(t *testing.T) {
	const (
		prefix = "!"
		val    = "xyz"
		expres = prefix + val
	)
	res, err := newAddPrefixStringActor(prefix).act(val)
	assert.Equal(t, res, expres)
	assert.NoError(t, err)
}
