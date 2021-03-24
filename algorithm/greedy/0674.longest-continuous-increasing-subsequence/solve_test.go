package problem0674

import "testing"

func TestSolve(t *testing.T) {
	t.Log(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	t.Log(findLengthOfLCIS([]int{2, 2, 2, 2, 2, 4}))
}
