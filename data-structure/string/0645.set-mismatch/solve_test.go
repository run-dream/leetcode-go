package problem0645

import "testing"

func TestSolve(t *testing.T) {
	t.Log(findErrorNums([]int{1, 2, 2, 4}))
	t.Log(findErrorNums([]int{1, 1}))
}
