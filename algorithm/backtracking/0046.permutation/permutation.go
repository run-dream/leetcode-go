package problem0046

func permute(nums []int) [][]int {
	result := [][]int{}
	permutation := []int{}
	visited := make([]bool, len(nums))
	dfs(permutation, visited, nums, &result)
	return result
}

func dfs(permutation []int, visited []bool, nums []int, result *[][]int) {
	if len(permutation) == len(nums) {
		*result = append(*result, append([]int{}, permutation...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		permutation = append(permutation, nums[i])
		dfs(permutation, visited, nums, result)
		permutation = permutation[0 : len(permutation)-1]
		visited[i] = false
	}
}
