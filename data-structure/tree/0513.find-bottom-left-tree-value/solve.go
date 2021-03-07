package problem0513

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	result := root.Val
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			node := queue[i]
			if i == 0 {
				result = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[cnt:]
	}
	return result
}
