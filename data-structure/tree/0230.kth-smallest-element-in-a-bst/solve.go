package problem0230

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	num := 0
	stack := []*TreeNode{}
	cur := root
	for num < k {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		if num == k-1 {
			return node.Val
		}
		num++
		stack = stack[:len(stack)-1]
		cur = node.Right
	}
	return -1
}
