package problem0328

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	isOdd := true
	cur := head.Next.Next
	oddCur := odd
	evenCur := even
	for cur != nil {
		next := cur.Next
		if isOdd {
			oddCur.Next = cur
			oddCur = cur
		} else {
			evenCur.Next = cur
			evenCur = cur
		}
		isOdd = !isOdd
		cur = next
	}
	oddCur.Next = even
	evenCur.Next = nil
	return odd
}
