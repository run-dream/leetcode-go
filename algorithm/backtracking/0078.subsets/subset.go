package problem0078

func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, subset, 0, i, &result)
	}
	return result
}

func backtrack(nums, subset []int, start, size int, result *[][]int) {
	if size == len(subset) {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtrack(nums, subset, i+1, size, result)
		subset = subset[:len(subset)-1]
	}
}
