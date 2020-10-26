package problem0257

import "testing"

func TestBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	t.Log(binaryTreePaths(root))
}
