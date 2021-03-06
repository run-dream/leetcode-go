package problem0111

import "testing"

func TestSolve(t *testing.T) {
	t.Log(minDepth(MakeTree([]int{3, 9, 20, 0, 0, 15, 7})))
	t.Log(minDepth(MakeTree([]int{2, 0, 3, 0, 4, 0, 5, 0, 6})))
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
