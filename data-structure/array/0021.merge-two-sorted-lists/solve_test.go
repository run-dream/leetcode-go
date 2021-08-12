package problem0021

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	l1 := makeList([]int{1, 2, 4})
	l2 := makeList([]int{1, 3, 44})
	traverseList(mergeTwoLists(l1, l2))
	traverseList(mergeTwoListsRecusive(l1, l2))
	traverseList(mergeTwoLists(makeList([]int{}), makeList([]int{})))
	traverseList(mergeTwoLists(makeList([]int{1}), makeList([]int{})))
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	p := head
	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{Val: nums[i]}
		p = p.Next
	}
	return head
}

func traverseList(node *ListNode) {
	fmt.Println("start")
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	fmt.Println("end")
}
