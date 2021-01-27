package problem0172

func trailingZeroes(n int) int {
	result := 0
	bound := n
	for bound > 0 {
		bound = bound / 5
		result += bound
	}
	return result
}
