package include

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type worker interface {
	doSomeWork() string
}

type workerImpl struct{}

func (wi *workerImpl) doSomeWork() string {
	return "w"
}

type workerDecorator struct {
	worker    worker
	workerPtr *worker
}

func (wi *workerDecorator) doSomeWork() string {
	return wi.doSomeWork()
}

func TestI(t *testing.T) {
	wd := workerDecorator{}
	assert.Nil(t, wd.worker)
	assert.Nil(t, *wd.workerPtr)
}
