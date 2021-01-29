package problem0496

import "testing"

func TestSolve(t *testing.T) {
	t.Log(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	t.Log(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
