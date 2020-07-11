package _struct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type worker interface {
	doSomeWork() string
}

type workerImpl struct{}

func (wi *workerImpl) doSomeWork() string {
	return "w"
}

type workerDecorator struct {
	worker worker
}

func (wi *workerDecorator) doSomeWork() string {
	return wi.doSomeWork()
}

func TestI(t *testing.T) {
	wd := workerDecorator{}
	assert.Nil(t, wd.worker)
}
