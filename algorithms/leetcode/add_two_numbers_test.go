package leetcode

import (
	"fmt"
	"testing"
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
//TODO -> TRY FIND BETTER SOLUTION
func Test_AddTwoNumbers(t *testing.T) {
	var ln1 *ListNode
	var ln2 *ListNode
	var res *ListNode

	ln1 = newListNode(2, 4, 3)
	ln2 = newListNode(5, 6, 4)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())

	ln1 = newListNode(0)
	ln2 = newListNode(1)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())

	ln1 = newListNode(5)
	ln2 = newListNode(5)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())

	ln1 = newListNode(8)
	ln2 = newListNode(10)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())

	ln1 = newListNode(0)
	ln2 = newListNode(7, 3)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())

	ln1 = newListNode(1)
	ln2 = newListNode(9, 9)
	res = addTwoNumbers(ln1, ln2)
	fmt.Println(res.StringVals())
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
