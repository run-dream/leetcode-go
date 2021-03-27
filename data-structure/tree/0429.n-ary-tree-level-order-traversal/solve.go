package problem0429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{root}
	nextQueue := []*Node{}
	for len(queue) > 0 {
		level := []int{}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			nextQueue = append(nextQueue, node.Children...)
		}
		queue = nextQueue
		result = append(result, level)
	}
	return result
}
