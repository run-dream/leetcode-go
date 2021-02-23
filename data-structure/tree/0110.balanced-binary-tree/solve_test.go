package problem0110

import "testing"

func TestSolve(t *testing.T) {
	root1 := MakeTree([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1})
	t.Log(isBalanced(root1))
	root2 := MakeTree([]int{3, 9, 20, 0, 0, 15, 7})
	t.Log(isBalanced(root2))
	root3 := MakeTree([]int{1, 2, 2, 3, 3, 0, 0, 4, 4})
	t.Log(isBalanced(root3))
}

func MakeTree(vals []int) *TreeNode {
	nodes := []*TreeNode{}
	for i := 0; i < len(vals); i++ {
		var node *TreeNode
		if vals[i] != 0 {
			node = &TreeNode{Val: vals[i]}
		}
		nodes = append(nodes, node)
	}
	for i := 0; i < len(vals)/2; i++ {
		if nodes[i] == nil {
			continue
		}
		nodes[i].Left = nodes[i*2+1]
		nodes[i].Right = nodes[i*2+2]
	}
	return nodes[0]
}
