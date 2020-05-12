package problem0452

import "testing"

func TestAssignCookie(t *testing.T) {
	t.Log((findMinArrowShots([][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}) == 2))
}
