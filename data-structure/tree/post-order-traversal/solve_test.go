package post_order_traversal

import "testing"

func TestSolve(t *testing.T) {
	root := &Node{val: 1}
	root.left = &Node{val: 2}
	root.right = &Node{val: 3}
	root.left.left = &Node{val: 4}
	root.left.right = &Node{val: 5}
	root.right.left = &Node{val: 6}
	root.right.right = &Node{val: 7}
	postOrder(root)
}
