package algorithms

import (
	"encoding/json"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	b, _ := json.Marshal(ln)
	return string(b)
}

func (ln *ListNode) StringVals() string {
	var b strings.Builder
	current := ln
	for {
		if current == nil {
			break
		}
		b.WriteString(strconv.Itoa(current.Val))
		current = current.Next
	}
	return b.String()
}

//noinspection SpellCheckingInspection
func NewListNode(vals ...int) *ListNode {
	if len(vals) < 0 {
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
