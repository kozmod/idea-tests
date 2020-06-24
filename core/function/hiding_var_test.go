package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const one = 1
const two = 2

func TestHideVar(t *testing.T) {
	x := one
	fmt.Println(x) // 1
	assert.Equal(t, one, x)

	{
		fmt.Println(x) // 1
		assert.Equal(t, one, x)
		x := 2
		fmt.Println(x) // 2
		assert.Equal(t, two, x)
	}

	fmt.Println(x) // 1
	assert.Equal(t, one, x)
}
