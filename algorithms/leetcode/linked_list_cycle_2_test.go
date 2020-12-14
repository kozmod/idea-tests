package leetcode

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
//Input: head = [3,2,0,-4], pos = 1
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
//
//Example 2:
//Input: head = [1,2], pos = 0
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
//
//Example 3:
//Input: head = [1], pos = -1
//Output: false
//Explanation: There is no cycle in the linked list.

func Test_Linked_List_Cycle_2(t *testing.T) {
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
	assert.Equal(t, head, detectCycle(head))

	assert.True(t, nil == detectCycle(nil))

	assert.True(t, nil == detectCycle(&ListNode{Val: 0}))

	assert.True(t, nil == detectCycle(&ListNode{Val: 0, Next: &ListNode{Val: 1}}))

	assert.Equal(t, head, detectCycleArray(head))

	assert.True(t, nil == detectCycleArray(nil))

	assert.True(t, nil == detectCycleArray(&ListNode{Val: 0}))

	assert.True(t, nil == detectCycleArray(&ListNode{Val: 0, Next: &ListNode{Val: 1}}))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	for depth := 0; next != nil; depth++ {
		if f := find2(head, next, depth); f != nil {
			return f
		}
		next = next.Next
	}
	return nil
}

func find2(head, search *ListNode, depth int) *ListNode {
	current := head
	for i := 0; depth > i; i++ {
		if current == search {
			return current
		}
		current = current.Next
	}
	return nil
}

func detectCycleArray(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	arr = append(arr, head)
	next := head.Next
	for depth := 0; next != nil; depth++ {
		for i := len(arr) - 1; i >= 0; i-- {
			tmp := arr[i]
			if tmp == next {
				return tmp
			}
		}
		arr = append(arr, next)
		next = next.Next
	}
	return nil
}
