package leetcode

import (
	. "github.com/kozmod/idea-tests/algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Palindrome Linked List
//Given a singly linked list, determine if it is a palindrome.
//
//Example 1:
//Input: 1->2
//Output: false
//
//Example 2:
//Input: 1->2->2->1
//Output: true
//
//Follow up:
//Could you do it in O(n) time and O(1) space?

func Test_Palindrome_Linked_List(t *testing.T) {
	list := NewLinkedListNode(1, 2, 2, 1)
	palindrome := isPalindrome(list)
	assert.True(t, palindrome)

	list = NewLinkedListNode(1, 2)
	palindrome = isPalindrome(list)
	assert.False(t, palindrome)

	list = NewLinkedListNode()
	palindrome = isPalindrome(list)
	assert.True(t, palindrome)

	list = NewLinkedListNode(1)
	palindrome = isPalindrome(list)
	assert.True(t, palindrome)

	list = NewLinkedListNode(1, 0, 1)
	palindrome = isPalindrome(list)
	assert.True(t, palindrome)
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var l []*ListNode

	current := head
	for current != nil {
		l = append(l, current)
		current = current.Next
	}
	if len(l) < 2 {
		return true
	}
	middle := len(l) / 2
	if len(l)%2 == 0 {
		for i, j := middle-1, middle; j < len(l); i, j = i-1, j+1 {
			left := l[i]
			right := l[j]
			if left.Val != right.Val {
				return false
			}
		}
	} else {
		for i, j := middle, middle; j < len(l); i, j = i-1, j+1 {
			if i < 0 {
				return false
			}
			left := l[i]
			right := l[j]
			if left.Val != right.Val {
				return false
			}
		}
	}
	return true
}
