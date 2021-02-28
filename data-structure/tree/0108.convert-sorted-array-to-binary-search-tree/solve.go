package problem0108

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return makeTree(nums, 0, len(nums)-1)
}

func makeTree(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = makeTree(nums, start, mid-1)
	node.Right = makeTree(nums, mid+1, end)
	return node
}
