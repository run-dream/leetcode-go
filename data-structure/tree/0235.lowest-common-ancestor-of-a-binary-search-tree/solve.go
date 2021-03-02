package problem0235

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	minVal, maxVal := p.Val, q.Val
	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}
	if minVal <= root.Val && maxVal >= root.Val {
		return root
	}
	if minVal <= root.Val && maxVal <= root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
