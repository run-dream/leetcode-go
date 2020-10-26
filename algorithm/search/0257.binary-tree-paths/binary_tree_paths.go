package problem0257

import "strconv"

import "strings"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var path []string
	dfs(root, path, &result)
	return result
}

func dfs(node *TreeNode, path []string, result *[]string) {
	if node == nil {
		return
	}
	path = append(path, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		*result = append(*result, strings.Join(path, "->"))
		return
	}
	dfs(node.Left, path, result)
	dfs(node.Right, path, result)
}
