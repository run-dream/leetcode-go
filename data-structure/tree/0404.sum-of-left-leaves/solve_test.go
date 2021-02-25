package problem0404

import "testing"

func TestSolve(t *testing.T) {
	t.Log(sumOfLeftLeaves(MakeTree([]int{3, 9, 20, 0, 0, 15, 7})))
	t.Log(sumOfLeftLeaves(MakeTree([]int{1})))
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
