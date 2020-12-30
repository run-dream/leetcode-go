package problem0037

func solveSudoku(board [][]byte) {
	solve(&board)
}

func solve(board *[][]byte) bool {
	x, y := findUnassigned(board)
	if x < 0 || y < 0 {
		return true
	}
	for i := 1; i <= 9; i++ {
		v := byte(i + '1' - 1)
		if isMatch(board, x, y, v) {
			(*board)[x][y] = v
			if solve(board) {
				return true
			}
			(*board)[x][y] = '.'
		}
	}
	return false
}

func findUnassigned(board *[][]byte) (x, y int) {
	x = -1
	y = -1
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if (*board)[i][j] == '.' {
				x = i
				y = j
				return
			}
		}
	}
	return
}

func isMatch(board *[][]byte, x, y int, v byte) bool {
	for i := 0; i < len((*board)[x]); i++ {
		if (*board)[x][i] == v {
			return false
		}
	}
	for i := 0; i < len(*board); i++ {
		if (*board)[i][y] == v {
			return false
		}
	}
	startRow := x - x%3
	startCol := y - y%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if (*board)[i][j] == v {
				return false
			}
		}
	}
	return true
}
