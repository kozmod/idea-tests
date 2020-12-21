package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestErrorFromContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	assert.Equal(t, ctx.Err(), context.Canceled)
	assert.Equal(t, ctx.Err(), context.Canceled)
	assert.Equal(t, <-ctx.Done(), struct{}{})

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Nanosecond)
	time.Sleep(2 * time.Nanosecond)
	assert.NotEqual(t, ctx.Err(), context.Canceled)
	assert.Equal(t, ctx.Err(), context.DeadlineExceeded)
	assert.Equal(t, <-ctx.Done(), struct{}{})
}
