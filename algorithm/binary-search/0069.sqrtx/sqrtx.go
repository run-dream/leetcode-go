package problem0069

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low := 0
	high := x
	for low <= high {
		mid := low + (high-low)/2
		sqrt := x / mid
		if mid == sqrt {
			return mid
		} else if mid < sqrt {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

func newtonSqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
