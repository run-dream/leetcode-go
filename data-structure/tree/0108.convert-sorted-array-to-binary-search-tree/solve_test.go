package problem0108

import "testing"

func TestSolve(t *testing.T) {
	printTree := func(root *TreeNode) {
		if root == nil {
			return
		}
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			node := queue[0]
			t.Log(node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	printTree(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
