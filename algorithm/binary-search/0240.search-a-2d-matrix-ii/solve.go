package problem0240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if target == matrix[row][col] {
			return true
		}
		if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}
	return false
}
