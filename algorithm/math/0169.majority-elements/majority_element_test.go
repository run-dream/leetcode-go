package problem0169

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Log(majorityElement([]int{1}))
	t.Log(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
