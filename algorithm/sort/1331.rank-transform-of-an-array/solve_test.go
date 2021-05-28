package problem1331

import "testing"

func TestSolve(t *testing.T) {
	t.Log(arrayRankTransform([]int{40, 10, 20, 30}))
	t.Log(arrayRankTransform([]int{100, 100, 100}))
	t.Log(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))

	t.Log(arrayRankTransformII([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
