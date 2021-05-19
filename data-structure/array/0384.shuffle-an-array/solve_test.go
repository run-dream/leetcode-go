package problem0384

import "testing"

func TestSolve(t *testing.T) {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	param1 := obj.Reset()
	t.Log(param1)
	param2 := obj.Shuffle()
	t.Log(param2)
}
