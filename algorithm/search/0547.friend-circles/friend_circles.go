package problem0547

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	found := make([]bool, len(M))
	num := 0
	for i := 0; i < len(M); i++ {
		if !found[i] {
			dfs(M, i, found)
			num++
		}
	}
	return num
}

func dfs(M [][]int, i int, found []bool) {
	found[i] = true
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && !found[j] {
			dfs(M, j, found)
		}
	}
}
