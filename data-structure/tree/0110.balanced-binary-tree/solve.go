package problem0110

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isBalance, _ := getDepth(root)
	return isBalance
}

func getDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, -1
	}
	leftBalance, leftDepth := getDepth(root.Left)
	if !leftBalance {
		return false, -1
	}
	rightBalance, rightDepth := getDepth(root.Right)
	if !rightBalance {
		return false, -1
	}
	if leftDepth-rightDepth > 1 || leftDepth-rightDepth < -1 {
		return false, -1
	}
	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}
	return true, depth + 1
}
