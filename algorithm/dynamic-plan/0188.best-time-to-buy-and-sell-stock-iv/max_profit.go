package problem0188

func maxProfit(k int, prices []int) int {
	n := len(prices)
	// 如果可交易次数超过一半 则可以把所以上涨的收益都获取到
	if k >= n/2 {
		maxPro := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxPro += prices[i] - prices[i-1]
			}
		}
		return maxPro
	}
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
	}
	// 对于每次交易
	for i := 1; i <= k; i++ {
		// 当前收益
		localMax := dp[i-1][0] - prices[0]
		for j := 1; j < n; j++ {
			// 在第j天是否进行交易
			dp[i][j] = max(dp[i][j-1], prices[j]+localMax)
			localMax = max(localMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
