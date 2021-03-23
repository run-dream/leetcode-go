package problem0598

func maxCount(m int, n int, ops [][]int) int {
	left := m
	right := n
	for i := 0; i < len(ops); i++ {
		left = min(left, ops[i][0])
		right = min(right, ops[i][1])
	}
	return left * right
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
