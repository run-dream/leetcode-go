package problem0540

func singleNonDuplicate(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 2
			} else {
				high = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return nums[low]
}
