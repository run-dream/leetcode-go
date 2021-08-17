package problem0024

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	PrintList(swapPairs(BuildList([]int{1, 2, 3, 4})))
	PrintList(swapPairs(BuildList([]int{1, 2, 3, 4, 5})))
	PrintList(swapPairsRecu(BuildList([]int{1, 2, 3, 4})))
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
