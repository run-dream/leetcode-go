package problem0442

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			result = append(result, nums[i])
		}
	}
	return result
}
