package problem0145

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := postorderTraversal(root.Left)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

func postorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{}
	cur := root
	var prev *TreeNode
	for cur != nil || len(stack) > 0 {
		// 访问左树
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 左右节点都访问过
			if cur.Right == nil || cur.Right == prev {
				result = append(result, cur.Val)
				prev = cur
				cur = nil
			} else {
				// 访问右树
				stack = append(stack, cur)
				cur = cur.Right
			}
		}
	}
	return result
}
