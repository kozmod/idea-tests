package linkedlist

import (
	"encoding/json"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	b, _ := json.Marshal(ln)
	return string(b)
}

//noinspection SpellCheckingInspection
func NewLinkedListNode(vals ...int) *ListNode {
	if vals == nil || len(vals) < 0 {
		return &ListNode{}
	}
	first := &ListNode{Val: vals[0]}
	previous := first
	for i := 1; i < len(vals); i++ {
		previous.Next = &ListNode{Val: vals[i]}
		previous = previous.Next
	}
	return first
}
