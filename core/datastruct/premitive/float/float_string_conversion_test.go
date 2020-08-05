package float

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatToString_strconv_FormatFloat(t *testing.T) {
	//noinspection ALL
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
