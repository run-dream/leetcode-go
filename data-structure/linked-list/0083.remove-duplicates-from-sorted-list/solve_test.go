package problem0083

import "testing"

func TestSolve(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 1}
	c := &ListNode{Val: 2}
	d := &ListNode{Val: 3}
	e := &ListNode{Val: 3}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	result := deleteDuplicates2(a)
	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}
