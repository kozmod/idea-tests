package easy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

/*
Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
1 <= n <= 45
*/

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		in  int
		exp int
	}{
		{
			in:  2,
			exp: 2,
		},
		{
			in:  3,
			exp: 3,
		},
		{
			in:  6,
			exp: 13,
		},
	}
	for i, testCase := range testCases {
		res := climbStairs(testCase.in)
		assert.True(t, reflect.DeepEqual(res, testCase.exp),
			fmt.Sprintf("expected:%v, actual:%v, testcase:%v \n", testCase.exp, res, i))
	}
}

func climbStairs(n int) int {
	memo := make([]int, n+1)
	return climbStairsReq(0, n, memo)
}

func climbStairsReq(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if last := memo[i]; last > 0 {
		return last
	}
	memo[i] = climbStairsReq(i+1, n, memo) + climbStairsReq(i+2, n, memo)
	return memo[i]
}
