package leetcode

import (
	"github.com/kozmod/idea-tests/algorithms"
	"reflect"
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
*/
//TODO -> MINE -> TRY FIND BETTER SOLUTION
func Test_AddTwoNumbers(t *testing.T) {
	var ln1 *algorithms.ListNode
	var ln2 *algorithms.ListNode
	var res *algorithms.ListNode

	ln1 = algorithms.NewListNode(2, 4, 3)
	ln2 = algorithms.NewListNode(5, 6, 4)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(7, 0, 8)))

	ln1 = algorithms.NewListNode(0)
	ln2 = algorithms.NewListNode(1)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(1)))

	ln1 = algorithms.NewListNode(5)
	ln2 = algorithms.NewListNode(5)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(0, 1)))

	ln1 = algorithms.NewListNode(8)
	ln2 = algorithms.NewListNode(10)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(8, 1)))

	ln1 = algorithms.NewListNode(0)
	ln2 = algorithms.NewListNode(7, 3)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(7, 3)))

	ln1 = algorithms.NewListNode(1)
	ln2 = algorithms.NewListNode(9, 9)
	res = addTwoNumbers(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(0, 0, 1)))
}

func addTwoNumbers(l1 *algorithms.ListNode, l2 *algorithms.ListNode) *algorithms.ListNode {
	var first *algorithms.ListNode
	var previous *algorithms.ListNode
	set := func(res int) {
		ln := &algorithms.ListNode{Val: res}
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
				previous.Next = &algorithms.ListNode{Val: x}
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
func Test_AddTwoNumbers_2(t *testing.T) {
	var ln1 *algorithms.ListNode
	var ln2 *algorithms.ListNode
	var res *algorithms.ListNode

	ln1 = algorithms.NewListNode(2, 4, 3)
	ln2 = algorithms.NewListNode(5, 6, 4)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(7, 0, 8)))

	ln1 = algorithms.NewListNode(0)
	ln2 = algorithms.NewListNode(1)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(1)))

	ln1 = algorithms.NewListNode(5)
	ln2 = algorithms.NewListNode(5)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(0, 1)))

	ln1 = algorithms.NewListNode(8)
	ln2 = algorithms.NewListNode(10)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(8, 1)))

	ln1 = algorithms.NewListNode(0)
	ln2 = algorithms.NewListNode(7, 3)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(7, 3)))

	ln1 = algorithms.NewListNode(1)
	ln2 = algorithms.NewListNode(9, 9)
	res = addTwoNumbers_2(ln1, ln2)
	assert.True(t, reflect.DeepEqual(res, algorithms.NewListNode(0, 0, 1)))
}

func addTwoNumbers_2(l1 *algorithms.ListNode, l2 *algorithms.ListNode) *algorithms.ListNode {
	dummyHead := &algorithms.ListNode{Val: 0}
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
			curr.Next = &algorithms.ListNode{Val: sum % 10}
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
			curr.Next = &algorithms.ListNode{Val: carry}
		}
	}
	return dummyHead.Next
}
