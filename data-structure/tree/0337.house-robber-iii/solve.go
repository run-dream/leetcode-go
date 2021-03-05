package problem0337

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	robRoot := root.Val
	if root.Left != nil {
		robRoot += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		robRoot += rob(root.Right.Left) + rob(root.Right.Right)
	}
	robChild := rob(root.Left) + rob(root.Right)
	return max(robRoot, robChild)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
