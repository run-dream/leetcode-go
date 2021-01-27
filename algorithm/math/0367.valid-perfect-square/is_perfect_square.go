package problem0367

func isPerfectSquare(num int) bool {
	left := 1
	right := num
	for left <= right {
		mid := (left + right) / 2
		value := mid * mid
		if value == num {
			return true
		} else if value <= num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
