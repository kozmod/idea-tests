package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMaxInSortedShiftedSlice(t *testing.T) {
	res, err := SearchShiftedMax([]int{6, 7, 8, 9, 1, 2, 3, 4, 5})
	assert.Equal(t, 9, res)
	assert.NoError(t, err)

	res, err = SearchShiftedMax([]int{9})
	assert.Equal(t, 9, res)
	assert.NoError(t, err)

	res, err = SearchShiftedMax([]int{})
	assert.Equal(t, -1, res)
	assert.Error(t, err)

	res, err = SearchShiftedMax(nil)
	assert.Equal(t, -1, res)
	assert.Error(t, err)
}
