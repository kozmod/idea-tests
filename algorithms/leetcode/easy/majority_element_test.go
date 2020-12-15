package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Majority Element
//Given an array of size n, find the majority element.
//The majority element is the element that appears more than ⌊ n/2 ⌋ times.
//
//You may assume that the array is non-empty and the majority element always exist in the array.
//
//Example 1:
//Input: [3,2,3]
//Output: 3
//
//Example 2:
//Input: [2,2,1,1,1,2,2]
//Output: 2

func Test_Majority_Element(t *testing.T) {
	res := majorityElement([]int{3, 2, 3})
	assert.Equal(t, 3, res)

	res = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	assert.Equal(t, 2, res)
}

func majorityElement(nums []int) int {
	m := make(map[int]int, len(nums))
	max := 0
	val := -1
	for _, num := range nums {
		i := m[num]
		i++
		m[num] = i
		if max < i {
			max = i
			val = num
		}
	}
	return val
}
