package problem0827

func largestIsland(grid [][]int) int {
	n := len(grid)
	color := 2
	count := map[int]int{}
	count[0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count[color] = mark(grid, i, j, color)
				color++
			}
		}
	}

	result := 0
	fullOne := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				result = max(result, getArea(grid, i, j, count))
				fullOne = false
			}
		}
	}
	if fullOne {
		return n * n
	}
	return result
}

func mark(grid [][]int, i, j, color int) int {
	n := len(grid)
	count := 0
	if i < 0 || i >= n || j < 0 || j >= n || grid[i][j] != 1 {
		return count
	}
	grid[i][j] = color
	return 1 + mark(grid, i+1, j, color) + mark(grid, i, j+1, color) +
		mark(grid, i-1, j, color) + mark(grid, i, j-1, color)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getArea(grid [][]int, i, j int, count map[int]int) int {
	n := len(grid)
	colors := map[int]bool{}
	directions := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	area := 1
	for _, direction := range directions {
		x := i + direction[0]
		y := j + direction[1]
		if x >= 0 && x < n && y >= 0 && y < n {
			color := grid[x][y]
			if !colors[color] {
				colors[color] = true
				area += count[color]
			}
		}
	}
	return area
}
