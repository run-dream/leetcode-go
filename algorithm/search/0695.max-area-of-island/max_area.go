package problem0695

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			current := dfs(grid, i, j)
			if current > max {
				max = current
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs(grid, i, j+1) + dfs(grid, i, j-1) + dfs(grid, i-1, j) + dfs(grid, i+1, j) + 1
}
