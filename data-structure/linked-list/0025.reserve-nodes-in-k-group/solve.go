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
	// 判断长度是不是够
	ptr := head
	for i := 0; i < k; i++ {
		ptr = ptr.Next
		if ptr == nil {
			return head
		}
	}

	// 反转head后k个元素
	var prev, next *ListNode
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	if cur != nil {
		head.Next = reverseKGroup(next, k)
	}

	return prev
}
