package problem0141

// ListNode 单连标
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && slow != nil {
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
