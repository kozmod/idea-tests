package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	list := NewListNode(1, 2, 3, 4, 5, 6, 7)

	rev := &ListNode{Val: list.Val}
	next := list.Next
	for {
		if next != nil {
			tmp := &ListNode{Val: rev.Val, Next: rev.Next}
			rev = &ListNode{Val: next.Val}
			rev.Next = tmp
			next = next.Next
		} else {
			break
		}
	}
	assert.True(t, reflect.DeepEqual(rev, NewListNode(7, 6, 5, 4, 3, 2, 1)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	b, _ := json.Marshal(ln)
	return string(b)
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
