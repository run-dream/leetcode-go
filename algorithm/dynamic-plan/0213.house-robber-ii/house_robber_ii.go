package problem0213

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robRange(nums []int, start, end int) int {
	if end == start {
		return nums[start]
	}
	if end == start+1 {
		return max(nums[start], nums[end])
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = nums[start+1]
	dp[start+2] = nums[start] + nums[start+2]
	for i := start + 3; i <= end; i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}
	return max(dp[end], dp[end-1])
}
