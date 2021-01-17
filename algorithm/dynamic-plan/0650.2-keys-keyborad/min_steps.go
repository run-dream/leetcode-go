package problem0650

func minSteps(n int) int {

	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := i - 1; j > 0; j-- {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
