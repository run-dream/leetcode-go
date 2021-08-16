package problem0011

import "testing"

func TestSolve(t *testing.T) {
	t.Log(maxArea([]int{1, 1}))
	t.Log(maxArea([]int{4, 3, 2, 1, 4}))
	t.Log(maxArea([]int{1, 2, 1}))
}
