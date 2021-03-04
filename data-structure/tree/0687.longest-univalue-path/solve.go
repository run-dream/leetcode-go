package problem0687

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	getPath(root, &result)
	return result
}

func getPath(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := getPath(root.Left, result)
	right := getPath(root.Right, result)
	if root.Left != nil && root.Left.Val == root.Val {
		left++
	} else {
		left = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right++
	} else {
		right = 0
	}
	*result = max(*result, left+right)
	return max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
