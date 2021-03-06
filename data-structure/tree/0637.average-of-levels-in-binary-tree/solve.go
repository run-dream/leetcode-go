package problem0637

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cnt := len(queue)
		sum := 0
		for i := 0; i < cnt; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[cnt:]
		result = append(result, float64(sum)/float64(cnt))
	}
	return result
}
