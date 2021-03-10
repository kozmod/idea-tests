package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAppendToSub(t *testing.T) {
	//goland:noinspection ALL
	print := func(prefix string, sub, origin []string) {
		fmt.Println(prefix, "\n",
			"sub ", sub,
			len(sub), cap(sub), "\n",
			"origin ", origin,
			len(origin), cap(origin))
	}

	origin := []string{"a", "b", "c", "d", "e"}
	sub := origin[1:3]
	print("1", sub, origin)
	assert.True(t, reflect.DeepEqual(sub, []string{"b", "c"}))
	assert.True(t, reflect.DeepEqual(origin, []string{"a", "b", "c", "d", "e"}))

	sub = append(sub, "&")
	sub = append(sub, "&")
	print("2", sub, origin)
	assert.True(t, reflect.DeepEqual(sub, []string{"b", "c", "&", "&"}))
	assert.True(t, reflect.DeepEqual(origin, []string{"a", "b", "c", "&", "&"}))

	sub = append(sub, "!")
	sub = append(sub, "!")
	print("3", sub, origin)
	assert.True(t, reflect.DeepEqual(sub, []string{"b", "c", "&", "&", "!", "!"}))
	assert.True(t, reflect.DeepEqual(origin, []string{"a", "b", "c", "&", "&"}))
}
