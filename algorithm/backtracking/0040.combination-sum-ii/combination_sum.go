package problem0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	visited := make([]bool, len(candidates))
	combination := []int{}
	sort.Ints(candidates)
	backtrack(candidates, target, 0, combination, visited, &result)
	return result
}

func backtrack(candidates []int, target int, start int, combination []int, visited []bool, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		if visited[i] {
			continue
		}
		if i != 0 && !visited[i-1] && candidates[i] == candidates[i-1] {
			continue
		}
		visited[i] = true
		combination = append(combination, candidates[i])
		backtrack(candidates, target-candidates[i], i+1, combination, visited, result)
		visited[i] = false
		combination = combination[:len(combination)-1]
	}
}
