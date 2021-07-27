package problem0733

import "testing"

func TestSolve(t *testing.T) {
	t.Log(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
}
