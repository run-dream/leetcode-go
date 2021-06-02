package problem0057

import "testing"

func TestSolve(t *testing.T) {
	t.Log(insert([][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}))
	t.Log(insert([][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}, []int{4, 8}))
	t.Log(insert([][]int{[]int{1, 5}}, []int{2, 3}))
	t.Log(insert([][]int{[]int{1, 5}}, []int{2, 7}))
	t.Log(insert([][]int{}, []int{2, 7}))
}
