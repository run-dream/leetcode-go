package problem0112

import "testing"

import "fmt"

func TestSolve(t *testing.T) {
	root1 := MakeTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1})
	t.Log(hasPathSum(root1, 22))

	root2 := MakeTree([]int{1, 2, 3})
	t.Log(hasPathSum(root2, 5))
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
		fmt.Print(i)
		nodes[i].Left = nodes[i*2+1]
		nodes[i].Right = nodes[i*2+2]
	}
	return nodes[0]
}
