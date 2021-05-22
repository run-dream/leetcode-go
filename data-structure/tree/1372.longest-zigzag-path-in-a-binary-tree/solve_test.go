package problem1372

import "testing"

func TestSolve(t *testing.T) {
	tree := &TreeNode{Val: 1}
	t.Log(longestZigZag(tree))
	tree.Left = &TreeNode{Val: 2}
	t.Log(longestZigZag(tree))
	tree.Left.Left = &TreeNode{Val: 3}
	t.Log(longestZigZag(tree))
}
