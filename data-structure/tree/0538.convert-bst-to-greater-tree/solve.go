package problem0538

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	current := root
	stack := []*TreeNode{}
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += node.Val
		current = node.Right
	}
	current = root
	prevVal := 0
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prevVal = node.Val
		node.Val = sum
		sum -= prevVal
		current = node.Right
	}
	return root
}
