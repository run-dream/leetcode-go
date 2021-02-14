package problem0160

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cursorA := headA
	cursorB := headB
	for cursorA != cursorB {
		if cursorA != nil {
			cursorA = cursorA.Next
		} else {
			cursorA = headB
		}
		if cursorB != nil {
			cursorB = cursorB.Next
		} else {
			cursorB = headA
		}
	}
	return cursorA
}
