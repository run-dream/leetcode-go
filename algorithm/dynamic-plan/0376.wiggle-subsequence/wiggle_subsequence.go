package problem0376

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 2 {
		if nums[0] != nums[1] {
			return 2
		}
		return 1
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	trend := getTrend(nums[1], nums[0])
	if trend == 0 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}
	ret := dp[1]
	for i := 2; i < n; i++ {
		newTrend := getTrend(nums[i], nums[i-1])
		if newTrend != trend && newTrend != 0 && (i != n-1 || trend != 0) {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
		if newTrend != 0 {
			trend = newTrend
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func getTrend(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}

func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}
