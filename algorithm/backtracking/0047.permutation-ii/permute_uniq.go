package problem0047

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	visited := make([]bool, len(nums))
	permutation := []int{}
	permute(permutation, visited, nums, &result)
	return result
}

func permute(permutation []int, visited []bool, nums []int, result *[][]int) {
	if len(permutation) == len(nums) {
		*result = append(*result, append([]int{}, permutation...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i >= 1 && nums[i] == nums[i-1] && visited[i-1] {
			continue
		}
		permutation = append(permutation, nums[i])
		visited[i] = true
		permute(permutation, visited, nums, result)
		visited[i] = false
		permutation = permutation[0 : len(permutation)-1]
	}
}
