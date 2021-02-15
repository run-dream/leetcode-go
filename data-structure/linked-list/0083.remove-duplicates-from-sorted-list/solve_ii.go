package problem0083

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates2(head.Next)
	if head.Next != nil && head.Val == head.Next.Val {
		return head.Next
	}
	return head
}
