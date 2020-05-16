package problem0655

func checkPossibility(nums []int) bool {
	modified := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if modified {
				return false
			}
			modified = true
			if i-2 >= 0 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}
	return true
}
