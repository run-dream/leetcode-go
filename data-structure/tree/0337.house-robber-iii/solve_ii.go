package problem0337

func rob2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rb, rc := robChild(root)
	return max(rb, rc)
}

func robChild(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lr, lc := robChild(root.Left)
	rr, rc := robChild(root.Right)

	return max(lr, lc) + max(rr, rc), root.Val + lr + rr
}
