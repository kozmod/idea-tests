package medium

import (
	. "github.com/kozmod/idea-tests/algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Linked List Cycle 2
//
//Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
//There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
//Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
//Notice that you should not modify the linked list.
//
//Example 1:
//  3 -> 2 -> 0 -> -4 -> nil
//       ^          |
//       |__________|
//Input: head = [3,2,0,-4], pos = 1
//Output: tail connects to node index 1
//Explanation: There is a cycle in the linked list, where tail connects to the second node.
//
//Example 2:
//  1 -> 2 -> nil
//  ^    |
//  |____|
//Input: head = [1,2], pos = 0
//Output: tail connects to node index 0
//Explanation: There is a cycle in the linked list, where tail connects to the first node.
//
//Example 3:
//  1 -> nil
//
//Input: head = [1], pos = -1
//Output: no cycle
//Explanation: There is no cycle in the linked list.
//
//Constraints:
//The number of the nodes in the list is in the range [0, 104].
//-105 <= Node.val <= 105
//pos is -1 or a valid index in the linked-list.
//
//Follow up: Can you solve it using O(1) (i.e. constant) memory?

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
