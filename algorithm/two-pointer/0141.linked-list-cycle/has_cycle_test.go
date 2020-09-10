package problem0141

import "testing"

func TestHasCyle(t *testing.T) {
	head := &ListNode{Val: 3, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 0, Next: nil}
	head.Next.Next.Next = &ListNode{Val: -4, Next: nil}
	head.Next.Next.Next.Next = head.Next
	t.Log(hasCycle(head))
}
