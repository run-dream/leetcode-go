package problem0142

import "testing"

func TestSolve(t *testing.T) {
	t.Log(detectCycle(buildCicleList([]int{3, 2, 0, -4}, 1)))
	t.Log(detectCycleTwoPointer(buildCicleList([]int{3, 2, 0, -4}, 1)))
}

func buildCicleList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	if pos < 0 {
		return head
	}
	ptr := head
	for pos > 0 {
		ptr = ptr.Next
		pos--
	}
	cur.Next = ptr
	return head
}
