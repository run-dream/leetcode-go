package problem0024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	newHead := head.Next
	for cur != nil && newHead != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
		if cur != nil && cur.Next != nil {
			tmp.Next = cur.Next
		}
	}
	return newHead
}

func swapPairsRecu(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}
