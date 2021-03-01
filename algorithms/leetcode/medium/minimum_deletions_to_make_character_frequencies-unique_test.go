package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

}

/*
Minimum Deletions to Make Character Frequencies Unique

A string s is called good if there are no two different characters in s that have the same frequency.

Given a string s, return the minimum number of characters you need to delete to make s good.

The frequency of a character in a string is the number of times it appears in the string.
For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.


Example 1:
Input: s = "aab"
Output: 0
Explanation: s is already good.

Example 2:
Input: s = "aaabbbcc"
Output: 2
Explanation: You can delete two 'b's resulting in the good string "aaabcc".
Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".

Example 3:
Input: s = "ceabaacb"
Output: 2
Explanation: You can delete both 'c's resulting in the good string "eabaab".
Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).


Constraints:
1 <= s.length <= 105
s contains only lowercase English letters.
*/

func Test_Minimum_Deletions_to_Make_Character_Frequencies_Unique(t *testing.T) {
	assert.Equal(t, 0, minDeletions("aab"))
	assert.Equal(t, 2, minDeletions("aaabbbcc"))
	assert.Equal(t, 2, minDeletions("ceabaacb"))
	assert.Equal(t, 2, minDeletions("bbcebab"))

}

func minDeletions(S string) int {
	quantity := make(map[rune]int)
	for _, val := range S {
		i := quantity[val]
		i++
		quantity[val] = i
	}
	unique := make(map[int]struct{}, len(quantity))
	res := 0
	for _, val := range quantity {
		for {
			_, ok := unique[val]
			if ok {
				res++
				if val = val - 1; val == 0 {
					break
				}
			} else {
				unique[val] = struct{}{}
				break
			}
		}
	}
	return res
}
