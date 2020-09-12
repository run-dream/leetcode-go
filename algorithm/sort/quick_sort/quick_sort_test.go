package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	quickSort(arr)
	t.Log(arr)
}
