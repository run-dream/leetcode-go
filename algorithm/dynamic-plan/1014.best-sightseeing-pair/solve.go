package problem1014

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	cur, res := 0, 0
	for i := 0; i < n; i++ {
		res = max(res, cur+values[i])
		cur = max(cur, values[i]) - 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
