package problem0160

import "testing"

func TestSolve(t *testing.T) {
	listA := &ListNode{Val: 4}
	listB := &ListNode{Val: 4}
	t.Log(getIntersectionNode(listA, listB) == nil)
}
