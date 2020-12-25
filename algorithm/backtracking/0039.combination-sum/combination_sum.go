package problem0039

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	combination := []int{}
	backtrack(candidates,combination, target, 0, &result)
	return result
}

func backtrack(candidates []int, combination []int, target int, start int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target{
			combination = append(combination, candidates[i])
			backtrack(candidates,combination, target - candidates[i], i ,result);
			combination = combination[:len(combination)-1]
		}
	}
}
