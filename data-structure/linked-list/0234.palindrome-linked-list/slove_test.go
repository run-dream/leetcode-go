package problem0234

import "testing"

func TestSolve(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	t.Log(isPalindrome(node1))
}
