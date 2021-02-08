package problem0448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}
