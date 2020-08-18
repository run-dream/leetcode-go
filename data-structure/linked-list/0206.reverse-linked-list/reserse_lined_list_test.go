package problem0206

import (
	"testing"
)

func TestReverse(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	var head = &ListNode{Val: 1, Next: nil}
	cursor := head
	for i := 1; i < len(vals); i++ {
		var tmp ListNode
		tmp.Val = vals[i]
		cursor.Next = &tmp
		cursor = cursor.Next
	}
	reverse := reverseList(head)
	for node := reverse; node != nil; node = node.Next {
		t.Log(node.Val)
	}
}
