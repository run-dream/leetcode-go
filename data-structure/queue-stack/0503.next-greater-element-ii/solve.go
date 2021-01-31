package problem0503

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = -1
	}
	for i := 0; i < len(nums)*2; i++ {
		index := i
		if i >= len(nums) {
			index = i - len(nums)
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[index] {
			result[stack[len(stack)-1]] = nums[index]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return result
}
