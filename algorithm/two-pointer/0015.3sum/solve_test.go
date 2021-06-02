package problem0015

import "testing"

func TestSolve(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{}))
	t.Log(threeSum([]int{0}))
	t.Log(threeSum([]int{0, 0, 0}))
}
