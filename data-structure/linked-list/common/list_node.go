package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
