package search

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMaxInSortedShiftedSlice(t *testing.T) {
	res, err := searchMax([]int{6, 7, 8, 9, 1, 2, 3, 4, 5})
	assert.Equal(t, 9, res)
	assert.NoError(t, err)

	res, err = searchMax([]int{9})
	assert.Equal(t, 9, res)
	assert.NoError(t, err)

	res, err = searchMax([]int{})
	assert.Equal(t, -1, res)
	assert.Error(t, err)

	res, err = searchMax(nil)
	assert.Equal(t, -1, res)
	assert.Error(t, err)
}

func searchMax(in []int) (int, error) {
	if len(in) == 1 {
		return in[0], nil
	}
	l := 0
	r := len(in) - 1
	for l <= r {
		if in[l] <= in[r] {
			return in[l-1], nil
		}
		mid := l + (r-l)/2
		next := (mid + 1) % len(in)
		prev := (mid + len(in) - 1) % len(in)
		if in[mid] <= in[next] && in[mid] <= in[prev] {
			return in[mid-1], nil
		}
		if in[mid] <= in[r] {
			r = mid - 1
		} else if in[mid] >= in[l] {
			l = mid + 1
		}
	}
	return -1, errors.New("array is empty")
}
