package problem0283

import "testing"

func TestSolve(t *testing.T) {
	array := []int{0, 1, 3, 0, 4, 12}
	moveZeroes(array)
	t.Log(array)
	array = []int{0, 1, 3, 0, 4, 12}
	moveZeroesTwoPointer(array)
	t.Log(array)
}
