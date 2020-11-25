package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Example 1:
//
//Input: s = "()"
//Output: true
//Example 2:
//
//Input: s = "()[]{}"
//Output: true
//Example 3:
//
//Input: s = "(]"
//Output: false
//Example 4:
//
//Input: s = "([)]"
//Output: false
//Example 5:
//
//Input: s = "{[]}"
//Output: true
func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		in  string
		res bool
	}{
		{
			in:  "()",
			res: true,
		},
		{
			in:  "()[]{}",
			res: true,
		},
		{
			in:  "{[]}",
			res: true,
		},
		{
			in:  "(]",
			res: false,
		},
		{
			in:  "([)]",
			res: false,
		},
		{
			in:  "(])",
			res: false,
		},
	}
	for i, testCase := range testCases {
		res := isValid(testCase.in)
		assert.Equal(t, testCase.res, res,
			fmt.Sprintf("expected: %v, got: %v, case: %d", testCase.res, res, i))
	}
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	p := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	n := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	tmp := make([]byte, 0, len(s))
	for _, c := range s {
		if len(tmp) < 1 {
			tmp = append(tmp, byte(c))
			continue
		}
		if _, ok := n[tmp[len(tmp)-1]]; ok {
			return false
		}
		bc := byte(c)
		if _, ok := p[bc]; ok {
			tmp = append(tmp, bc)
			continue
		}

		neg := n[bc]
		if neg == tmp[len(tmp)-1] {
			tmp = tmp[:len(tmp)-1]
		} else {
			return false
		}
	}
	if len(tmp) > 0 {
		return false
	}
	return true
}
