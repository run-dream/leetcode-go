package problem1764

import "testing"

func TestSolve(t *testing.T) {
	t.Log(canChoose([][]int{
		[]int{1, -1, -1},
		[]int{3, -2, 0},
	}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}))
	t.Log(canChoose([][]int{
		[]int{10, -2},
		[]int{1, 2, 3, 4},
	}, []int{1, 2, 3, 4, 10, -2}))
	t.Log(canChoose([][]int{
		[]int{1, 2, 3},
		[]int{3, 4},
	}, []int{7, 7, 1, 2, 3, 3, 4, 7, 7}))
}
