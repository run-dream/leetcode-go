package problem0827

import "testing"

func TestSolve(t *testing.T) {
	t.Log(largestIsland([][]int{
		[]int{1, 0, 0, 1, 1},
		[]int{1, 0, 0, 1, 0},
		[]int{1, 1, 1, 1, 1},
		[]int{1, 1, 1, 0, 1},
		[]int{0, 0, 0, 1, 0},
	}))
}
