package bubble_sort

import (
	"testing"
)

func TestSolve(t *testing.T) {
	arr := []int{1, 3, 2, 4}
	t.Log(arr)
	bubbleSort(arr)
	t.Log(arr)
}
