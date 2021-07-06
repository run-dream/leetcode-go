package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Push(3)
	t.Log(heap.Top())
	heap.Push(4)
	t.Log(heap.Top())
	heap.Push(2)
	t.Log(heap.Top())
	heap.Pop()
	t.Log(heap.Top())
	heap.Pop()
	t.Log(heap.Top())
}
