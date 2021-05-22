package problem1372

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	dfs(root.Left, true, 1)
	dfs(root.Right, false, 1)
	return max
}

func dfs(node *TreeNode, isLeft bool, count int) {
	if node == nil {
		if count-1 > max {
			max = count - 1
		}
		return
	}
	if isLeft {
		dfs(node.Right, false, count+1)
		dfs(node.Left, true, 1)
	} else {
		dfs(node.Left, true, count+1)
		dfs(node.Right, false, 1)
	}
}
