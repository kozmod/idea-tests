package leetcode

import (
	"fmt"
	. "github.com/kozmod/idea-tests/algorithms/linkedlist"
	"testing"
)

//Reverse Linked List

//Example:
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
//Follow up:
//A linked list can be reversed either iteratively or recursively. Could you implement both?

func TestReverseLinkedList(t *testing.T) {
	list := NewLinkedListNode(1, 2, 3)
	res := reverseList(list)
	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	rev := &ListNode{Val: head.Val}
	next := head.Next
	for next != nil {
		tmp := &ListNode{Val: rev.Val, Next: rev.Next}
		rev = &ListNode{Val: next.Val, Next: tmp}
		next = next.Next
	}
	return rev
}
