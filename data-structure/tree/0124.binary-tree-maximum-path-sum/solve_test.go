package problem0124

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Log(maxPathSum(MakeTree([]interface{}{-10, 9, 20, nil, nil, 15, 7})))
}

func MakeTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := []*TreeNode{}
	for i := 0; i < len(vals); i++ {
		var node *TreeNode
		if vals[i] != nil {
			node = &TreeNode{Val: vals[i].(int)}
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
