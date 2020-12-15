package easy

import (
	. "github.com/kozmod/idea-tests/algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Linked List Cycle
//
//Given head, the head of a linked list, determine if the linked list has a cycle in it.
//There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
//Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
//Return true if there is a cycle in the linked list. Otherwise, return false.
//
//Example 1:
//  3 -> 2 -> 0 -> -4 -> nil
//       ^          |
//       |__________|
//Input: head = [3,2,0,-4], pos = 1
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
//
//Example 2:
//  1 -> 2 -> nil
//  ^    |
//  |____|
//Input: head = [1,2], pos = 0
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
//
//Example 3:
//  1 -> nil
//
//Input: head = [1], pos = -1
//Output: false
//Explanation: There is no cycle in the linked list.

func Test_Linked_List_Cycle(t *testing.T) {
	head := &ListNode{Val: 0}
	last := &ListNode{
		Val:  2,
		Next: head,
	}
	mid := &ListNode{
		Val:  1,
		Next: last,
	}
	head.Next = mid
	assert.True(t, hasCycle(head))

	assert.False(t, hasCycle(nil))

	assert.False(t, hasCycle(&ListNode{Val: 0}))

	assert.False(t, hasCycle(&ListNode{Val: 0, Next: &ListNode{Val: 1}}))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycleReq(head *ListNode) bool {
	if head == nil {
		return false
	}
	next := head.Next
	for depth := 0; next != nil; depth++ {
		if find(head, next, depth) {
			return true
		}
		next = next.Next
	}
	return false
}

func hasCycleArray(head *ListNode) bool {
	if head == nil {
		return false
	}
	arr := make([]*ListNode, 0, 0)
	next := head.Next
	arr = append(arr, head)
	for next != nil {
		for _, node := range arr {
			if node == next {
				return true
			}
		}
		arr = append(arr, next)
		next = next.Next
	}
	return false
}

func find(head, search *ListNode, depth int) bool {
	current := head
	for i := 0; depth > i; i++ {
		if current == search {
			return true
		}

		current = current.Next
	}
	return false
}
