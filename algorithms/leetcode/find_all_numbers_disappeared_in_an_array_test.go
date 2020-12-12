package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//Find All Numbers Disappeared in an Array
//Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
//Find all the elements of [1, n] inclusive that do not appear in this array.
//
//Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//Example:
//Input:
//[4,3,2,7,8,2,3,1]
//Output:
//[5,6]

func TestFind_All_Numbers_Disappeared_in_an_Array__Stupid(t *testing.T) {
	testCase := []struct {
		in  []int
		exp []int
	}{
		{
			[]int{1, 1},
			[]int{2},
		},
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
	}

	for i, tc := range testCase {
		res := findDisappearedNumbers(tc.in)
		assert.True(t, reflect.DeepEqual(tc.exp, res), fmt.Sprintf("testcase:%d, ext: %v, res:%v", i, tc.exp, res))
	}

}

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	var res []int
	for i := 1; i <= len(nums); i++ {
		_, ok := m[i]
		if !ok {
			res = append(res, i)
		}
	}
	return res
}
