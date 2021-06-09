package problem0106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	end := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			end = i
			break
		}
	}
	root.Left = buildTree(inorder[:end], postorder[:end])
	root.Right = buildTree(inorder[end+1:], postorder[end:len(postorder)-1])
	return root
}
