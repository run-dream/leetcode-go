package problem0025

import "testing"

func TestSolve(t *testing.T) {
	t.Log(reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 2))
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}
