package problem1014

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	left, res := values[n-1]-n-1, 0
	for i := n - 2; i >= 0; i-- {
		left = max(left, values[i+1]-i-1)
		res = max(res, left+values[i]+i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
