package sort_test

import (
	"testing"

	"github.com/kozmod/idea-tests/algorithms/classic/sort"
	"github.com/stretchr/testify/assert"
)

func TestCountSortParallel(t *testing.T) {
	testCases := []struct {
		Name     string
		Numbers  []int32
		Expected []int32
		Range    int32
	}{
		{
			Name:     "All the numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9, 6, 3, 8, 7, 2, 5},
			Expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Range:    int32(10),
		},
		{
			Name:     "3 numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9},
			Expected: []int32{1, 4, 9},
			Range:    int32(10),
		},
	}

	for _, tc := range testCases {
		tc := tc // We run our tests twice one with this line & one without
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := sort.CountSort(tc.Numbers, tc.Range)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, res))
		})
	}
}
