package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCancelContext_CancelMoreThenOne(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	assert.NoError(t, ctx.Err())
	cancel()
	assert.Error(t, ctx.Err())
	cancel()
	assert.Error(t, ctx.Err())
}

func TestCancelContext_CancelByTimeout(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	assert.NoError(t, ctx.Err())
	time.Sleep(2 * time.Second)
	assert.Error(t, ctx.Err())
	assert.Error(t, ctx.Err())
	cancel()
	assert.Error(t, ctx.Err())
}
