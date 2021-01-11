package problem0416

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
		if nums[i] < half {
			dp[i][nums[i]] = true
		} else if nums[i] == half {
			return true
		}
	}
	for i := 1; i < len(nums); i++ {
		// 这里逆序处理，否则会导致覆盖
		for j := half; j >= 0; j-- {
			dp[i][j] = dp[i-1][j]
			if !dp[i][j] && j >= nums[i] {
				dp[i][j] = dp[i-1][j-nums[i]]
			}
		}
		if dp[i][half] {
			return true
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	dp := make([]bool, half+1)
	if nums[0] <= half {
		dp[nums[0]] = true
	}
	if dp[half] {
		return true
	}
	for i := 1; i < len(nums); i++ {
		for j := half; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[half]
}
