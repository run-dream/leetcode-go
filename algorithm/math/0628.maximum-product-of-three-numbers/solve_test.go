package problem0628

import "testing"

func TestSolve(t *testing.T) {
	t.Log(maximumProduct([]int{1, 2, 3, 4}))
	t.Log(maximumProduct([]int{-1, -2, -3, -4}))
	t.Log(maximumProduct([]int{1, 2, 3, -3, -4}))
}
