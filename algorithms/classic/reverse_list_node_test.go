package classic

import (
	"fmt"
	"github.com/kozmod/idea-tests/algorithms"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	list := algorithms.NewListNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	rev := &algorithms.ListNode{Val: list.Val}
	next := list.Next
	for {
		if next != nil {
			tmp := &algorithms.ListNode{Val: rev.Val, Next: rev.Next}
			rev = &algorithms.ListNode{Val: next.Val}
			rev.Next = tmp
			next = next.Next
		} else {
			break
		}
	}
	fmt.Println(rev)
}
