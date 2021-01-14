package problem0322

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	// 用 amount + 1 表示不可达
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if i >= coins[j] && dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] < amount+1 {
		return dp[amount]
	}
	return -1
}
