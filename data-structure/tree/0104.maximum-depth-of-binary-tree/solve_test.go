package problem0104

import "testing"

func TestSolve(t *testing.T) {
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 20}
	node4 := &TreeNode{Val: 15}
	node5 := &TreeNode{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	t.Log(maxDepth(node1))
}
