package problem0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	p1, p2, p3 := l1.Next, l2, l1
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p3.Next = p2
			break
		}
		if p2 == nil {
			p3.Next = p1
			break
		}
		if p1.Val < p2.Val {
			p3.Next = p1
			p3 = p3.Next
			p1 = p1.Next
		} else {
			p3.Next = p2
			p3 = p3.Next
			p2 = p2.Next
		}
	}
	return l1
}

func mergeTwoListsRecusive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecusive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecusive(l1, l2.Next)
	return l2
}
