package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Generate Parentheses

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]


Constraints:
1 <= n <= 8
*/

func Test_Generate_Parentheses(t *testing.T) {
	assert.Equal(t,
		[]string(nil),
		generateParenthesis(0))
	assert.Equal(t,
		[]string{"()"},
		generateParenthesis(1))
	assert.Equal(t,
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string
	if n > 0 {
		generate(&res, "", 0, 0, n)
	}
	return res

}

func generate(res *[]string, current string, open, closed, max int) {
	if len(current) == max*2 {
		*res = append(*res, current)
	}
	if open < max {
		generate(res, current+"(", open+1, closed, max)
	}
	if closed < open {
		generate(res, current+")", open, closed+1, max)
	}
}
