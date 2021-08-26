package problem0124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32

	oneSide(root, &result)

	return result
}

func oneSide(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}
	left := max(0, oneSide(node.Left, result))
	right := max(0, oneSide(node.Right, result))

	*result = max(left+right+node.Val, *result)

	return max(left, right) + node.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
