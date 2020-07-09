package context_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type key string

const (
	UidKey key = "uid"
	anyKey     = "anyKey"
	anyVal     = "anyKey"
)

func GetUid(ctx context.Context) string {
	if uid, ok := ctx.Value(UidKey).(string); !ok {
		// WTF???
		return ""
	} else {
		return uid
	}
}

func PutUid(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, UidKey, uid)
}

func TestContextWithParameters(t *testing.T) {
	uid := "1234-5678-xxx-yxz-901"
	ctx := PutUid(context.Background(), uid)
	ctx = context.WithValue(ctx, anyKey, anyVal)
	worker(ctx, "some_msg")
	assert.Equal(t, uid, GetUid(ctx))
	assert.Equal(t, anyVal, ctx.Value(anyKey).(string))
}

func worker(ctx context.Context, s string) {
	fmt.Println("work:", s, ",", "ctx:", GetUid(ctx))
}
