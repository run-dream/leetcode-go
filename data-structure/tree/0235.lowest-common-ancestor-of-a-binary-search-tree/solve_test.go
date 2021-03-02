package problem0235

import "testing"

func TestSolve(t *testing.T) {
	t.Log(lowestCommonAncestor(MakeTree([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}), &TreeNode{Val: 2}, &TreeNode{Val: 8}).Val)
	t.Log(lowestCommonAncestor(MakeTree([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}), &TreeNode{Val: 2}, &TreeNode{Val: 4}).Val)
	t.Log(lowestCommonAncestor(MakeTree([]int{2, 1, -1}), &TreeNode{Val: 2}, &TreeNode{Val: 1}).Val)
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
