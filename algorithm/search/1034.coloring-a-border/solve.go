package problem1034

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) == 0 {
		return grid
	}
	oldColor := grid[r0][c0]
	height := len(grid)
	width := len(grid[0])
	bfs(grid, r0, c0, height, width, oldColor, color)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == -1 {
				grid[i][j] = oldColor
			} else if grid[i][j] == -2 {
				grid[i][j] = color
			}
		}
	}
	return grid
}

func bfs(grid [][]int, r0, c0, height, width, oldColor, newColor int) {
	if !isPosLegal(r0, c0, height, width) {
		return
	}
	if grid[r0][c0] < 0 {
		return
	}
	if grid[r0][c0] != oldColor {
		return
	}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bound := false
	for _, direction := range directions {
		r1 := direction[0] + r0
		c1 := direction[1] + c0
		if !isPosLegal(r1, c1, height, width) {
			bound = true
		} else {
			nextColor := grid[r1][c1]
			if nextColor != oldColor && nextColor > 0 {
				bound = true
			}
		}
	}
	if bound {
		grid[r0][c0] = -2
	} else {
		// 用-1来标记已经访问
		grid[r0][c0] = -1
	}
	for _, direction := range directions {
		r1 := direction[0] + r0
		c1 := direction[1] + c0
		bfs(grid, r1, c1, height, width, oldColor, newColor)
	}
}

func isPosLegal(r, c, h, w int) bool {
	return r >= 0 && r < h && c >= 0 && c < w
}
