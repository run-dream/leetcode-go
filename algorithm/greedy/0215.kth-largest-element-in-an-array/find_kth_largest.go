package problem0215

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	// 第k大的元素在从小到大排列的len(nums)-k位
	target := len(nums) - k
	// 实际上在循环内必定能返回 但是这里没有break不能些 for{ 所以用循环过程中 left一定小于或等于right来标示
	for left <= right {
		// 快排分区
		pivot := partition(nums, left, right)
		if pivot == target {
			return nums[target]
		} else if pivot > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return nums[target]
}

func partition(nums []int, left, right int) int {
	// 取数组最右的元素作为基准
	pivot := nums[right]
	// i记录已经处理的区间 [left...i]
	i := left
	// 遍历 数组 [left, right-1]
	for j := left; j <= right-1; j++ {
		// 如果nums[j]比基准小 置于已处理数组的最右端
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将已处理区间的下一个元素和基准交换位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
