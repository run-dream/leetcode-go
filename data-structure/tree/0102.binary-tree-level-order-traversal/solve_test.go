package problem0102

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Log(levelOrder(MakeTree([]int{3, 9, 20, -1, -1, 15, 7})))
}

func MakeTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
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
		fmt.Print(i)
		nodes[i].Left = nodes[i*2+1]
		nodes[i].Right = nodes[i*2+2]
	}
	return nodes[0]
}
