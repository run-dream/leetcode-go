package problem0441

func arrangeCoins(n int) int {
	low := 1
	high := n
	for low <= high{
		mid := low + (high - low) / 2
		value := (1+mid) * mid / 2
		if value == n {
			return mid
		}else if value < n {
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return high
}