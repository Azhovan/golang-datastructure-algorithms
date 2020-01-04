package heap

import (
	"testing"
)

func TestHeap_Push(t *testing.T) {
	heap := NewMaxHeap()
	heap.Push(3)
	heap.Push(1)
	heap.Push(2)
	heap.Push(5)
	heap.Push(0)

	p1 := heap.Pop()
	p2 := heap.Pop()

	if p1 != 5 || p2 != 3 {
		t.Errorf("wrong pop, expect 5 and 3, got %d, %d %v", p1, p2, heap.List())
	}
}
