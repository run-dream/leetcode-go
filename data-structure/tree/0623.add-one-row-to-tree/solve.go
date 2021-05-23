package problem0623

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val}
		newRoot.Left = root
		return newRoot
	}
	h := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			node := queue[i]
			if h == depth-1 {
				newLeftChild := &TreeNode{Val: val}
				newLeftChild.Left = node.Left
				newRightChild := &TreeNode{Val: val}
				newRightChild.Right = node.Right
				node.Left = newLeftChild
				node.Right = newRightChild
			} else {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		queue = queue[levelLen:]
		if h == depth-1 {
			break
		}
		h++
	}
	return root
}
