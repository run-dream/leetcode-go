package problem0530

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minDif := 0xFFFF
	var prev *TreeNode
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && node != nil {
			diff := abs(node.Val - prev.Val)
			minDif = min(minDif, diff)
		}
		if minDif == 0 {
			return minDif
		}
		prev = node
		cur = node.Right
	}
	return minDif
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
