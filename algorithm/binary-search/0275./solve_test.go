package problem0275

import "testing"

func TestSolve(t *testing.T) {
	t.Log(hIndex([]int{0, 1, 3, 5, 6}))
	t.Log(hIndex([]int{1, 2, 100}))
	t.Log(hIndex([]int{100}))
	t.Log(hIndex([]int{0, 0, 4, 4}))
}
