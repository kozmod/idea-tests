package heap

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func TestHeap_Unsafe(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])

	assert.Equal(t, 1, heap.Pop(h))
	assert.Equal(t, 2, heap.Pop(h))
	assert.Equal(t, 3, heap.Pop(h))
	assert.Equal(t, 5, heap.Pop(h))

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap_Safe(t *testing.T) {
	h := newSafeIntHeap(2, 1, 5)
	h.Push(3)

	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 5, h.Pop())

}

type SafeIntHeap struct {
	heap *IntHeap
}

func newSafeIntHeap(vars ...int) *SafeIntHeap {
	h := IntHeap{}
	h = append(h, vars...)
	heap.Init(&h)
	return &SafeIntHeap{&h}
}

func (h *SafeIntHeap) Push(x int) {
	*h.heap = append(*h.heap, x)
}

func (h *SafeIntHeap) Pop() interface{} {
	x := heap.Pop(h.heap)
	return x.(int)
}

func (h *SafeIntHeap) Len() int {
	return h.heap.Len()
}
