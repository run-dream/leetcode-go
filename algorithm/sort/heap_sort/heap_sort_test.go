package heap_sort

import "testing"

func TestHeapSort(t *testing.T) {
	data := []int{1, 3, 4, 2, 5, 7, 6, 9, 8}
	Sort(data)
	t.Log(data)
}
