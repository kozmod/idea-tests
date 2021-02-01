package medium

import (
	"fmt"
	. "github.com/kozmod/idea-tests/algorithms/linkedlist"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

/*
Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

func Test_AddTwoNumbers(t *testing.T) {
	type testcase struct {
		inA *ListNode
		inB *ListNode
		exp *ListNode
	}
	testcases := []testcase{
		{
			inA: NewLinkedListNode(2, 4, 3),
			inB: NewLinkedListNode(5, 6, 4),
			exp: NewLinkedListNode(7, 0, 8),
		},
		{
			inA: NewLinkedListNode(0),
			inB: NewLinkedListNode(1),
			exp: NewLinkedListNode(1),
		},
		{
			inA: NewLinkedListNode(5),
			inB: NewLinkedListNode(5),
			exp: NewLinkedListNode(0, 1),
		},
		{
			inA: NewLinkedListNode(8),
			inB: NewLinkedListNode(10),
			exp: NewLinkedListNode(8, 1),
		},
		{
			inA: NewLinkedListNode(0),
			inB: NewLinkedListNode(7, 3),
			exp: NewLinkedListNode(7, 3),
		},
		{
			inA: NewLinkedListNode(1),
			inB: NewLinkedListNode(9, 9),
			exp: NewLinkedListNode(0, 0, 1),
		},
	}
	assertTest := func(f func(l1 *ListNode, l2 *ListNode) *ListNode) {
		for i, testcase := range testcases {
			res := f(testcase.inA, testcase.inB)
			assert.True(t, reflect.DeepEqual(res, testcase.exp),
				fmt.Sprintf("expected:%v, actual:%v, testcase:%v, func:%s \n",
					testcase.exp, res, i, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()))
		}
	}
	assertTest(addTwoNumbers)
	assertTest(addTwoNumbers_2)
	assertTest(addTwoNumbers_3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var first *ListNode
	var previous *ListNode
	set := func(res int) {
		ln := &ListNode{Val: res}
		if first == nil {
			first = ln
			previous = ln
		} else {
			previous.Next = ln
			previous = ln
		}
	}
	c1 := l1
	c2 := l2
	x := 0
	for {
		if c1 != nil && c2 != nil {
			res := c1.Val + c2.Val + x
			res, x = resultAndOver(res)

			c1 = c1.Next
			c2 = c2.Next
			set(res)
		} else if c1 != nil {
			if x > 0 {
				res := c1.Val + x
				res, x = resultAndOver(res)
				c1 = c1.Next
				set(res)
			} else {
				previous.Next = c1
				break
			}
		} else if c2 != nil {
			if x > 0 {
				res := c2.Val + x
				res, x = resultAndOver(res)
				c2 = c2.Next
				set(res)
			} else {
				previous.Next = c2
				break
			}
		} else {
			if x > 0 {
				previous.Next = &ListNode{Val: x}
			}
			break
		}
	}
	return first
}

func resultAndOver(res int) (int, int) {
	x := 0
	if res > 9 {
		if i := res % 10; i != 0 {
			x = res - i
			res = i
		} else {
			x = res / 10
			res = 0
		}
		if x > 9 {
			x = x / 10
		}
	}
	return res, x
}

//GOOD
func addTwoNumbers_2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	p := l1
	q := l2
	curr := dummyHead
	carry := 0
	for {
		if p != nil || q != nil {
			var x int
			var y int
			if p != nil {
				x = p.Val
			} else {
				x = 0
			}
			if q != nil {
				y = q.Val
			} else {
				y = 0
			}
			sum := carry + x + y
			carry = sum / 10
			curr.Next = &ListNode{Val: sum % 10}
			curr = curr.Next
			if p != nil {
				p = p.Next
			}
			if q != nil {
				q = q.Next
			}
		} else {
			break
		}
		if carry > 0 {
			curr.Next = &ListNode{Val: carry}
		}
	}
	return dummyHead.Next
}

//BEST
func addTwoNumbers_3(l1 *ListNode, l2 *ListNode) *ListNode {
	left := l1
	right := l2
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	sum := 0
	for {
		if left != nil && right != nil {
			sum = sum + left.Val + right.Val
			left = left.Next
			right = right.Next
		} else if left != nil {
			sum = sum + left.Val
			left = left.Next
		} else if right != nil {
			sum = sum + right.Val
			right = right.Next
		} else {
			if sum > 0 {
				current.Next = &ListNode{Val: sum}
				break
			}
			break
		}
		tmp := &ListNode{Val: sum}
		if tmp.Val >= 10 {
			tmp.Val = tmp.Val - 10
			sum = 1
		} else {
			sum = 0
		}
		current.Next = tmp
		current = current.Next
	}
	return dummyHead.Next
}
