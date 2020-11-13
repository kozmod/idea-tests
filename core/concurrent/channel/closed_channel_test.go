package channel

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPanic_WHenWriteToClosedChannel(t *testing.T) {
	defer func() {
		p := recover()
		assert.NotNil(t, p)
	}()
	ch := make(chan bool)
	close(ch)
	ch <- true
}

func TestPanic_WhenTryCloseClosedChannel(t *testing.T) {
	defer func() {
		p := recover()
		assert.NotNil(t, p)
	}()
	ch := make(chan bool)
	close(ch)
	close(ch)
}

func TestFalse_WhenTryReadFromClosedChannel(t *testing.T) {
	ch := make(chan bool)
	close(ch)
	val, ok := <-ch
	assert.False(t, val)
	assert.False(t, ok)
}

func TestTryReadFromClosedBufferedChannel(t *testing.T) {
	const val = "some value"
	ch := make(chan string, 2)
	ch <- val
	ch <- val
	close(ch)
	time.Sleep(1 * time.Second)
	res, ok := <-ch
	assert.Equal(t, val, res)
	assert.NotEqual(t, "", res)
	assert.True(t, ok)

	res, ok = <-ch
	assert.Equal(t, val, res)
	assert.NotEqual(t, "", res)
	assert.True(t, ok)

	res, ok = <-ch
	assert.NotEqual(t, val, res)
	assert.Equal(t, "", res)
	assert.False(t, ok)

	res, ok = <-ch
	assert.NotEqual(t, val, res)
	assert.Equal(t, "", res)
	assert.False(t, ok)
}
