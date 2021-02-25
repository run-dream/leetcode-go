package problem0404

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root, 0, false)
}

func dfs(root *TreeNode, sum int, fromLeft bool) int {
	if root == nil {
		return sum
	}
	if root.Left == nil && root.Right == nil {
		if fromLeft {
			return sum + root.Val
		}
		return sum
	}
	sum = dfs(root.Left, sum, true)
	sum = dfs(root.Right, sum, false)
	return sum
}
