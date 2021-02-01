package problem0594

import "testing"

func TestSolve(t *testing.T) {
	t.Log(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	t.Log(findLHS([]int{1, 2, 3, 4}))
	t.Log(findLHS([]int{1, 1, 1, 1}))
}
