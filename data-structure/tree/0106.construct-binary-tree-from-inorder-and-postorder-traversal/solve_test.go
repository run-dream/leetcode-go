package problem0106

import "testing"

func TestSolve(t *testing.T) {
	t.Log(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	t.Log(buildTree([]int{-1}, []int{-1}))
	t.Log(buildTree([]int{}, []int{}))
}
