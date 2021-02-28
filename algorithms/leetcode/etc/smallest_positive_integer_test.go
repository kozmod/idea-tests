package etc

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*

Smallest positive integer

Write a function:

class Solution { public int solution(int[] A); }

that, given an array A of N integers,
returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/
func TestSmallestPositiveInteger(t *testing.T) {
	assert.Equal(t, 4, Solution([]int{1, 2, 3}))
	assert.Equal(t, 5, Solution([]int{1, 3, 6, 4, 1, 2}))
	assert.Equal(t, 1, Solution([]int{-1, -4}))

	assert.Equal(t, 4, Solution_2([]int{1, 2, 3}))
	assert.Equal(t, 5, Solution_2([]int{1, 3, 6, 4, 1, 2}))
	assert.Equal(t, 1, Solution_2([]int{-1, -4}))
}

func Solution(A []int) int {
	sort.Ints(A)
	max := 0
	for _, val := range A {
		if max < val {
			if max+1 == val {
				max = val
			} else {
				break
			}
		}
	}
	max = max + 1
	return max
}

func Solution_2(A []int) int {
	m := make(map[int]struct{})
	for _, val := range A {
		m[val] = struct{}{}
	}
	max := 0
	for _, val := range A {
		_, ok := m[val+1]
		if !ok {
			if val > 0 {
				if max == 0 {
					max = val
				}
				if val < max {
					max = val
				}
			}
		}
	}
	max = max + 1
	return max
}
