package block

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestVarVisibility(t *testing.T) {
	x := 0
	{
		x := 1
		assert.Equal(t, x, 1)

	}
	assert.Equal(t, x, 0)
	{
		x = 2
		assert.Equal(t, x, 2)
	}
	assert.Equal(t, x, 2)
}
