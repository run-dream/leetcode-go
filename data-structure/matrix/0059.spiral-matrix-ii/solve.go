package problem0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	directions := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	x, y, i, d := 0, 0, 1, 0
	matrix[x][y] = i
	for i < n*n {
		direction := directions[d]
		x += direction[0]
		y += direction[1]
		if x >= 0 && x < n && y >= 0 && y < n && matrix[x][y] == 0 {
			i++
			matrix[x][y] = i
		} else {
			x -= direction[0]
			y -= direction[1]
			d = (d + 1) % 4
		}
	}
	return matrix
}
