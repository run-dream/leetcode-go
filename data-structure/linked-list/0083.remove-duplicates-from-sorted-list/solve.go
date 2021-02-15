package problem0083

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cur := head.Next
	for cur != nil {
		if prev.Val != cur.Val {
			prev.Next = cur
			prev = cur
		}
		cur = cur.Next
	}
	if prev.Next != nil && prev.Next.Val == prev.Val {
		prev.Next = nil
	}
	return head
}
