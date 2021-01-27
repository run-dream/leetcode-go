package problem0494

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	dp := make([][]int, n)
	maxVal := 0
	for _, val := range nums {
		maxVal += val
	}
	if S < -maxVal || S > maxVal {
		return 0
	}
	size := 2*maxVal + 1
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, size)
	}
	dp[0][maxVal-nums[0]] = 1
	dp[0][maxVal+nums[0]] = dp[0][maxVal+nums[0]] + 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < size; j++ {
			sum := 0
			if j+nums[i] < size {
				sum += dp[i-1][j+nums[i]]
			}
			if j-nums[i] >= 0 {
				sum += dp[i-1][j-nums[i]]
			}
			dp[i][j] = sum
		}
	}
	return dp[n-1][S+maxVal]
}
