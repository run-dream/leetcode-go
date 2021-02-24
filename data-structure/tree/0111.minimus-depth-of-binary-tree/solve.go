package problem0111

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return dfs(root.Right) + 1
	}
	if root.Right == nil {
		return dfs(root.Left) + 1
	}
	return min(dfs(root.Right), dfs(root.Left)) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
