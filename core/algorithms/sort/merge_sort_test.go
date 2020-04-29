package sort_test

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{1, 8, 99, 0, -3, 7}
	exp := []int{-3, 0, 1, 7, 8, 99}
	res := sort(s)
	assert.Equal(t, exp, res)
}

func sort(array []int) []int {
	inLen := len(array)
	if inLen < 2 {
		return array
	}
	left := sort(array[:inLen/2])
	right := sort(array[inLen/2 : inLen])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
