package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFor1(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	var calls int
	l := func(i []int) int {
		length := len(i)
		calls++
		return length
	}
	for i := 0; i < l(s); i++ {

	}
	assert.Equal(t, calls, len(s)+1)
}
