package problem0300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	ret := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		tmp := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				tmp = max(tmp, dp[j]+1)
			}
		}
		dp[i] = tmp
		ret = max(dp[i], ret)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
