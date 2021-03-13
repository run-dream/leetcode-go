package problem0501

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	maxTimes := 0
	curTimes := 0
	stack := []*TreeNode{}
	cur := root
	var prev *TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == nil || prev.Val != node.Val {
			curTimes = 1
		} else {
			curTimes++
		}
		if curTimes > maxTimes {
			result = []int{node.Val}
			maxTimes = curTimes
		} else if curTimes == maxTimes {
			result = append(result, node.Val)
		}
		prev = node
		cur = node.Right
	}
	return result
}
