package problem0287

func findDuplicate(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		cnt := 0
		for _, val := range nums {
			if val <= mid {
				cnt++
			}
		}
		if cnt > mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
