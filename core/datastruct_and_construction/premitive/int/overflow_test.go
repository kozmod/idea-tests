package int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntOverflow(t *testing.T) {
	var num int32 = 2147483647
	assert.Equal(t, int32(-2147483648), num+1)
}
