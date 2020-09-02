package problem0633

import "math"

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}
	return false
}
