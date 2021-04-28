package problem0275

func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		value := citations[mid]
		if value == n-mid {
			return n - mid
		} else if value > n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return n - left
}
