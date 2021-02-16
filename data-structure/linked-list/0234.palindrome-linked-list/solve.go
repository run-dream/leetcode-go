package problem0234

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	odd := false
	// 偶数个让slow指向下一个节点
	if fast != nil {
		slow = slow.Next
		odd = true
	}
	// 此时 slow为后半段的起始节点
	// 将前半部分截断，并翻转
	cur := head
	var prev *ListNode
	for cur != slow {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 奇数个，比较时跳过后半段的第一个元素
	if !odd {
		slow = slow.Next
	}
	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}

	return true
}
