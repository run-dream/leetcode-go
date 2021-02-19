package problem0617

import "testing"

func TestSolve(t *testing.T) {
	a1 := &TreeNode{Val: 1}
	a2 := &TreeNode{Val: 3}
	a3 := &TreeNode{Val: 2}
	a4 := &TreeNode{Val: 5}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	b1 := &TreeNode{Val: 2}
	b2 := &TreeNode{Val: 1}
	b3 := &TreeNode{Val: 3}
	b4 := &TreeNode{Val: 4}
	b5 := &TreeNode{Val: 7}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b4.Right = b5
	m := mergeTrees(a1, b1)
	queue := []*TreeNode{}
	queue = append(queue, m)
	for len(queue) != 0 {
		head := queue[0]
		t.Log(head.Val)
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
}
