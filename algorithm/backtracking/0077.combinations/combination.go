package problem0077

func combine(n int, k int) [][]int {
	combination := []int{}
	result := [][]int{}
	backtrack(1, n, k, combination, &result)
	return result
}

func backtrack(start, n, k int, combination []int, result *[][]int) {
	if k == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}
	for i := start; i <= n-k+1; i++ {
		combination = append(combination, i)
		backtrack(i+1, n, k-1, combination, result)
		combination = combination[:len(combination)-1]
	}
}
