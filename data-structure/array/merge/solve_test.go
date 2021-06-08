package offer_merge

import "testing"

func TestSolve(t *testing.T) {
	a := []int{1, 2, 5, -1, -1}
	merge(a, []int{3, 8}, 3)
	t.Log(a)
}
