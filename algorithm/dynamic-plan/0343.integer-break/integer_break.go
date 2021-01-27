package problem0343

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			tmp := (i - j) * j
			if dp[i-j]*j > tmp {
				tmp = dp[i-j] * j
			}
			if tmp > dp[i] {
				dp[i] = tmp
			}
		}
	}
	return dp[n]
}
