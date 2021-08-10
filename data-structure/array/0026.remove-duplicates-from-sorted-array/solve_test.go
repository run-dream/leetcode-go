package problem0026

import "testing"

func TestSolve(t *testing.T) {
	t.Log(removeDuplicates([]int{1, 2}))
	t.Log(removeDuplicates([]int{1, 1, 2}))
	t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
