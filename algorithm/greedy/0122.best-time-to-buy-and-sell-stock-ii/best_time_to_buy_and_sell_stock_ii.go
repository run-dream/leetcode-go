package problem0122

func maxProfit(prices []int) int {
	profix := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profix += prices[i] - prices[i-1]
		}
	}
	return profix
}
