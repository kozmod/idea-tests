package sort_test

import (
	"testing"

	"github.com/kozmod/idea-tests/algorithms/classic/sort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	s := []int{1, 8, 99, 0, -3, 7}
	exp := []int{-3, 0, 1, 7, 8, 99}
	res := sort.MergeSort(s)
	assert.Equal(t, exp, res)
}

func TestMergeSortParallel(t *testing.T) {
	testCases := []struct {
		Name     string
		Numbers  []int
		Expected []int
	}{
		{
			Name:     "All the numbers in the range [1-9]",
			Numbers:  []int{4, 1, 9, 6, 3, 8, 7, 2, 5},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			Name:     "3 numbers in the range [1-9]",
			Numbers:  []int{4, 1, 9},
			Expected: []int{1, 4, 9},
		},
		{
			Name:     "big positive and negative numbers",
			Numbers:  []int{1, 8, 99, 0, -3, 7, -77},
			Expected: []int{-77, -3, 0, 1, 7, 8, 99},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := sort.MergeSort(tc.Numbers)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, res), tc.Name)
		})
	}
}
