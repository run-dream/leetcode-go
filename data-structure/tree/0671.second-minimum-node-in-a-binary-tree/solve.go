package problem0671

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		return -1
	}
	leftVal := root.Left.Val
	rightVal := root.Right.Val
	if leftVal == root.Val {
		leftVal = findSecondMinimumValue(root.Left)
	}
	if rightVal == root.Val {
		rightVal = findSecondMinimumValue(root.Right)
	}
	if leftVal != -1 && rightVal != -1 {
		return min(leftVal, rightVal)
	}
	if leftVal != -1 {
		return leftVal
	}
	return rightVal
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
