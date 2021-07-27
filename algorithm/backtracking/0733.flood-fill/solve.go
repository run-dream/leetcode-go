package problem0733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return image
	}
	width := len(image)
	height := len(image[0])
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	bfs(&image, sr, sc, width, height, oldColor, newColor)
	return image
}

func bfs(image *[][]int, sr, sc, width, height, oldColor, newColor int) {
	if sr < 0 || sr >= width || sc < 0 || sc >= height {
		return
	}
	if (*image)[sr][sc] != oldColor {
		return
	}
	(*image)[sr][sc] = newColor
	directions := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for _, direction := range directions {
		bfs(image, sr+direction[0], sc+direction[1], width, height, oldColor, newColor)
	}
}
