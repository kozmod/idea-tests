package leetcode

import (
	"fmt"
	. "github.com/kozmod/idea-tests/algorithms"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//Merge two sorted linked lists and return it as a new sorted list.
//The new list should be made by splicing together the nodes of the first two lists.
//
//Input: l1 = [1,2,4], l2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//
//Input: l1 = [], l2 = []
//Output: []
//
//Input: l1 = [], l2 = [0]
//Output: [0]
func TestMergeTwoSortedSlices(t *testing.T) {
	testCases := []struct {
		l1  []int
		l2  []int
		res []int
	}{
		{
			l1:  []int{1, 2, 4},
			l2:  []int{1, 3, 4},
			res: []int{1, 1, 3, 4, 4, 4},
		},
		{
			l1:  []int{4},
			l2:  []int{},
			res: []int{4},
		},
		{
			l1:  []int{},
			l2:  []int{9},
			res: []int{9},
		},
		{
			l1:  []int{},
			l2:  []int{},
			res: []int{},
		},
	}
	for i, testCase := range testCases {
		res := mergeTwoSortedSlices(testCase.l1, testCase.l2)
		assert.Equal(t, testCase.res, res,
			fmt.Sprintf("expected: %v, got: %v, case: %d", testCase.res, res, i))
	}
}

func mergeTwoSortedSlices(l1, l2 []int) []int {
	len1 := len(l1)
	if len1 < 1 {
		return l2
	}
	len2 := len(l2)
	if len2 < 1 {
		return l1
	}
	res := make([]int, 0, len(l1)+len(l2))
	for _, v1 := range l1 {
		for j, v2 := range l2 {
			if v1 < v2 {
				res = append(res, v2)
			} else if v1 == v2 {
				res = append(res, v1, v2)
				l2 = l2[j:]
				break
			}
		}
	}
	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{
			l1:  NewListNode(1, 2, 3),
			l2:  NewListNode(1, 2, 4),
			res: NewListNode(1, 1, 2, 2, 3, 4),
		},
		{
			l1:  nil,
			l2:  NewListNode(0),
			res: NewListNode(0),
		},
		{
			l1:  NewListNode(1),
			l2:  nil,
			res: NewListNode(1),
		},
	}
	for i, testCase := range testCases {
		res := mergeTwoLists(testCase.l1, testCase.l2)
		assert.True(t, reflect.DeepEqual(testCase.res, res),
			fmt.Sprintf("expected: %v,\n got: %v,\n case: %d\n", testCase.res, res, i))
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	current1 := l1
	current2 := l2
	var tmp *ListNode
	var res *ListNode
	//goland:noinspection ALL
	for current1 != nil && current2 != nil {
		if res == nil {
			if current1.Val < current2.Val {
				res = &ListNode{Val: current1.Val}
				current1 = current1.Next
			} else {
				res = &ListNode{Val: current2.Val}
				current2 = current2.Next
			}
			tmp = res
			continue
		}
		if current1.Val < current2.Val && tmp != nil {
			tmp.Next = &ListNode{Val: current1.Val}
			current1 = current1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = &ListNode{Val: current2.Val}
			current2 = current2.Next
			tmp = tmp.Next
		}
	}
	if current1 == nil && tmp != nil {
		tmp.Next = current2
	} else if current1 == nil && tmp == nil {
		res = current2
	}
	if current2 == nil && tmp != nil {
		tmp.Next = current1
	} else if current2 == nil && tmp == nil {
		res = current1
	}
	return res
}
