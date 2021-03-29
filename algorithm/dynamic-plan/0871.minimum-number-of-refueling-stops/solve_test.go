package problem0871

import "testing"

func TestSplve(t *testing.T) {
	t.Log(minRefuelStopsII(1, 1, [][]int{}))
	t.Log(minRefuelStopsII(100, 1, [][]int{[]int{10, 100}}))
	t.Log(minRefuelStopsII(100, 10, [][]int{[]int{10, 60}, []int{20, 30}, []int{30, 30}, []int{60, 40}}))
}
