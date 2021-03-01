package problem0543

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	depth(root, &result)
	return result
}

func depth(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, result)
	rightDepth := depth(root.Right, result)
	*result = max(*result, leftDepth+rightDepth)
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
