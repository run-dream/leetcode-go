package problem0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	visited := make([]bool, len(nums))
	subset := []int{}
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, visited, subset, 0, i, &result)
	}
	return result
}

func backtrack(nums []int, visited []bool, subset []int, start int, size int, result *[][]int) {
	if len(subset) == size {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := start; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i != 0 && nums[i-1] == nums[i] && !visited[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		visited[i] = true
		backtrack(nums, visited, subset, i+1, size, result)
		visited[i] = false
		subset = subset[:len(subset)-1]
	}
}
