package problem1780

func checkPowersOfThree(n int) bool {
	for n > 0 {
		digit := n % 3
		if digit == 2 {
			return false
		}
		n /= 3
	}
	return true
}
