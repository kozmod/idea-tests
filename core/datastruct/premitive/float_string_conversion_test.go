package premitive

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFloatToString_strconv_FormatFloat(t *testing.T) {
	var f float64 = 1.12345678901
	s := strconv.FormatFloat(f, 'f', -1, 64)
	assert.Equal(t, s, "1.12345678901")

	s = strconv.FormatFloat(f, 'f', 1, 64)
	assert.Equal(t, s, "1.1")

	defer func() {
		if p := recover(); p != nil {
			assert.NotNil(t, p)
		}
	}()
	s = strconv.FormatFloat(f, 'f', -1, 16)
	assert.Equal(t, s, "1.12345678901")
}