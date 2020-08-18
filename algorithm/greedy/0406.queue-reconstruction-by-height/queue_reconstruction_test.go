package problem0406

import "testing"

func TestAssignQueueReconstruction(t *testing.T) {
	t.Log(reconstructQueue([][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}))
}
