package problem0025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 先判断输入是否合法
	if head == nil || k < 2 {
		return head
	}
	// 找到第k个元素
	ptr := head
	for i := 0; i < k-1; i++ {
		ptr = ptr.Next
		if ptr == nil {
			return head
		}
	}

	nextHead := ptr.Next
	ptr.Next = nil

	// 反转head后k个元素
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	head.Next = reverseKGroup(nextHead, k)

	return prev
}
