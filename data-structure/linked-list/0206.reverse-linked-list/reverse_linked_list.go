package problem0206

// ListNode singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	cur := head
	for cur != nil {
		// 要处理的下个指针
		next := cur.Next
		// 反转指针
		cur.Next = newHead
		// 新的队头
		newHead = cur
		// 更新指针
		cur = next
	}
	return newHead
}

func reverseListRecu(head *ListNode) *ListNode {
	// 定义基准情况
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecu(head.Next)
	// head.Next 是 newHead 的最后一个节点
	head.Next.Next = head
	// 清空 head.Next 指针
	head.Next = nil
	return newHead
}
