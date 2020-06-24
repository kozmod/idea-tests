package ref

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

type some struct {
	Val string
}

func TestEquality(t *testing.T) {
	a := some{"A"}
	b := some{"A"}
	c := &a
	d := &b

	equals := a == b
	fmt.Println(equals)
	assert.Equal(t, true, equals)

	noequals := c == d
	fmt.Println(noequals)
	assert.Equal(t, false, noequals)
}
