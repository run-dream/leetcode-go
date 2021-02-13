package problem0378

import "testing"

func TestSolve(t *testing.T) {
	t.Log(kthSmallest([][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8))
}
