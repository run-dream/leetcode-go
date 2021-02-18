package problem0226

import "testing"

func TestSolve(t *testing.T) {
	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 9}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	root := invertTree(node1)
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		head := queue[0]
		t.Log(head.Val)
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
}
