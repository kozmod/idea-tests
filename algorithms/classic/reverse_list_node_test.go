package classic

import (
	"fmt"
	"github.com/kozmod/idea-tests/algorithms/linkedlist"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	list := linkedlist.NewListNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

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
	fmt.Println(rev)
}
