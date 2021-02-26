package problem0669

import "testing"

func TestSolve(t *testing.T) {
	t.Log(trimBST(MakeTree([]int{1, 0, 2}), 1, 2))
	t.Log(trimBST(MakeTree([]int{4, 0, 4, -1, 2, -1, -1, 1}), 1, 3))
}

func MakeTree(vals []int) *TreeNode {
	nodes := []*TreeNode{}
	for i := 0; i < len(vals); i++ {
		var node *TreeNode
		if vals[i] != -1 {
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
