package problem0079

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int, visited [][]bool) bool {
	if k >= len(word) {
		return true
	}
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	if visited[i][j] {
		return false
	}
	visited[i][j] = true

	k = k + 1
	directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	for _, direction := range directions {
		ik := i + direction[0]
		jk := j + direction[1]
		if dfs(board, word, ik, jk, k, visited) {
			return true
		}
	}
	visited[i][j] = false
	return false
}
