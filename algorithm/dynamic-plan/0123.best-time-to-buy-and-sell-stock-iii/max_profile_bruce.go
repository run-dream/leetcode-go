package problem0123

func maxProfitBruce(prices []int) int {
	n := len(prices)
	profits1 := getProfits1(prices)
	profits2 := getProfits2(prices)
	result := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			result = maxBruce(result, profits1[i]+profits2[j])
		}
	}
	result = maxBruce(result, profits1[n-1])
	result = maxBruce(result, profits2[0])
	return result
}

func maxBruce(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minBruce(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getProfits1(prices []int) []int {
	mins := make([]int, len(prices))
	profits := make([]int, len(prices))
	mins[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		mins[i] = minBruce(mins[i-1], prices[i])
		// 第i天卖出的最大收益
		profits[i] = prices[i] - mins[i-1]
	}
	return profits
}

func getProfits2(prices []int) []int {
	maxs := make([]int, len(prices))
	profits := make([]int, len(prices))
	n := len(prices)
	maxs[n-1] = prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxs[i] = maxBruce(maxs[i+1], prices[i])
		//　第i天买入的最大收益
		profits[i] = maxs[i+1] - prices[i]
	}
	return profits
}
