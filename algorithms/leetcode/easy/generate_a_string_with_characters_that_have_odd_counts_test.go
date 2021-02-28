package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 Generate a String With Characters That Have Odd Counts

Given an integer n, return a string with n characters such that each character in such string occurs an odd number of times.
The returned string must contain only lowercase English letters. If there are multiples valid strings, return any of them.

Example 1:
Input: n = 4
Output: "pppz"
Explanation: "pppz" is a valid string since the character 'p' occurs three times and the character 'z' occurs once.
Note that there are many other valid strings such as "ohhh" and "love".

Example 2:
Input: n = 2
Output: "xy"
Explanation: "xy" is a valid string since the characters 'x' and 'y' occur once.
Note that there are many other valid strings such as "ag" and "ur".

Example 3:
Input: n = 7
Output: "holasss"


Constraints:
1 <= n <= 500
*/

func Test_Generate_a_String_With_Characters_That_Have_Odd_Counts(t *testing.T) {
	assert.Equal(t, "a", generateTheString(1))
	assert.Equal(t, "ab", generateTheString(2))
	assert.Equal(t, "aaa", generateTheString(3))
	assert.Equal(t, "aaab", generateTheString(4))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func generateTheString(N int) string {
	res := make([]rune, N, N)
	var mainRune = letterRunes[0]
	if N%2 == 0 {
		res[N-1] = letterRunes[1]
		N--
	}
	for i := 0; i < N; i++ {
		res[i] = mainRune
	}
	return string(res)
}
