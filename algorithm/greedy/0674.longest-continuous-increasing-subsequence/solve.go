package problem0674

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
			result = max(cur, result)
		} else {
			cur = 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
