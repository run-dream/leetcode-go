package problem0628

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[0] > 0 {
		return nums[n-1] * nums[n-2] * nums[n-3]
	}
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
