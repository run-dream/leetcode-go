package problem0278

func isBadVersion(version int) bool {
	return version >= 4
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
