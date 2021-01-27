package problem0123

func maxProfit(prices []int) int {
	firstBuy := (int)(-10E5 - 1)
	firstSell := 0
	secondBuy := (int)(-10E5 - 1)
	secondSell := 0
	for _, curPrice := range prices {
		if firstBuy < -curPrice {
			firstBuy = -curPrice
		}
		if firstSell < firstBuy+curPrice {
			firstSell = firstBuy + curPrice
		}
		if secondBuy < firstSell-curPrice {
			secondBuy = firstSell - curPrice
		}
		if secondSell < secondBuy+curPrice {
			secondSell = secondBuy + curPrice
		}
	}
	return secondSell
}
