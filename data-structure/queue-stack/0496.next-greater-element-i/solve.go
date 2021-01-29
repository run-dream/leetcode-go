package problem0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	dict := map[int]int{}
	for _, num2 := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num2 {
			dict[stack[len(stack)-1]] = num2
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num2)
	}
	result := []int{}
	for _, num1 := range nums1 {
		val, found := dict[num1]
		if found {
			result = append(result, val)
		} else {
			result = append(result, -1)
		}
	}
	return result
}
