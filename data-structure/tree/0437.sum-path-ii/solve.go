package problem0437

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	result, _ := count(root, sum)
	return result
}

func count(root *TreeNode, sum int) (int, map[int]int) {
	rootCount := map[int]int{}
	if root == nil {
		return 0, rootCount
	}
	sumLeft, leftCount := count(root.Left, sum)
	sumRight, rightCount := count(root.Right, sum)
	for leftKey, leftValue := range leftCount {
		rootCount[leftKey+root.Val] = leftValue
	}
	for rightKey, rightValue := range rightCount {
		rootCount[rightKey+root.Val] += rightValue
	}
	rootCount[root.Val]++
	return sumLeft + sumRight + rootCount[sum], rootCount
}
