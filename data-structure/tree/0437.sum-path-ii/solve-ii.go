package problem0437

func pathSumII(root *TreeNode, sum int) int {
	preSums := make(map[int]int)
	preSums[0] = 1
	return dfs(root, sum, 0, preSums)
}

func dfs(root *TreeNode, sum, curSum int, preSums map[int]int) int {
	if root == nil {
		return 0
	}
	curSum += root.Val
	// 如果curSum-sum存在说明存在路径可达
	total := preSums[curSum-sum]
	preSums[curSum]++
	total += dfs(root.Left, sum, curSum, preSums)
	total += dfs(root.Right, sum, curSum, preSums)
	preSums[curSum]--
	return total
}
