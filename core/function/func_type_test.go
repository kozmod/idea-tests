package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type stringActor interface {
	act(string) (string, error)
}

type stringActorImpl func(string) (string, error)

func (act stringActorImpl) act(s string) (string, error) {
	return act(s)
}

func newAddPrefixStringActor(prefix string) stringActorImpl {
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
	var actor stringActor = newAddPrefixStringActor(prefix)
	res, err := actor.act(val)
	assert.Equal(t, res, expres)
	assert.NoError(t, err)
}
