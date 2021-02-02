package classic

import (
	"github.com/kozmod/idea-tests/algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	list := linkedlist.NewLinkedListNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	rev := &linkedlist.ListNode{Val: list.Val}
	next := list.Next
	for {
		if next != nil {
			tmp := &linkedlist.ListNode{Val: rev.Val, Next: rev.Next}
			rev = &linkedlist.ListNode{Val: next.Val}
			rev.Next = tmp
			next = next.Next
		} else {
			break
		}
	}
	assert.True(t, reflect.DeepEqual(rev, linkedlist.NewLinkedListNode(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)))
}
