package problem0653

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	nums := inOrder(root)
	if len(nums) == 1 {
		return false
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		}
		if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrder(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		cur = node.Right
	}
	return result
}
