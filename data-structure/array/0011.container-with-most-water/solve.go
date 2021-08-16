package problem0011

func maxArea(height []int) int {
	size := len(height)
	if size <= 1 {
		return 0
	}
	i, j := 0, size-1
	water := 0
	for i < j {
		water = max(water, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
