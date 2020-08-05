package include

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	err  = errors.New("0")
	errA = errors.New("A")
)

type IntA interface {
	doWork(string) error
	doWorkA(string) error
}

type IntB interface {
	doWork(string) error
}

type SomeImpl struct{}

func (impl *SomeImpl) doWork(s string) error {
	return err
}

func (impl *SomeImpl) doWorkA(s string) error {
	return errA
}

func TestDucklingMultiInterfaceImpl(t *testing.T) {
	var a IntA = &SomeImpl{}
	var b IntB = a
	assert.NotNil(t, a)
	assert.NotNil(t, b)
	assert.NotNil(t, &a)
	assert.NotNil(t, &b)
	assert.Equal(t, a.doWork(""), err)
	assert.Equal(t, b.doWork(""), err)
	assert.NotNil(t, a.doWorkA(""), errA)
}
