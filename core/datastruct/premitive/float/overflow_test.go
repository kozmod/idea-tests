package float

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFloatOverflow(t *testing.T) {
	var f float32 = 16777216
	f2 := f + 1
	assert.Equal(t, f, f+1)
	assert.Equal(t, f, f2)
}
