package utils_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsSecond(t *testing.T) {
	sec := "1"
	res, err := AsSeconds(sec)
	assert.NoError(t, err)
	assert.Equal(t, time.Duration(1)*time.Second, res)
}

func TestAsSecondMaxInt32WhenError(t *testing.T) {
	sec := "xsxs"
	res, err := AsSeconds(sec)
	assert.Error(t, err)
	assert.Equal(t, time.Duration(math.MaxInt32)*time.Second, res)
}
