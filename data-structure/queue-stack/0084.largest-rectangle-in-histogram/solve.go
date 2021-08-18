package problem0084

func largestRectangleArea(heights []int) int {
	size := len(heights)
	if size == 0 {
		return 0
	}
	// 第 i 个位置的值对应向左找到的小于它的第一个值
	left := make([]int, size)
	left[0] = -1
	for i := 1; i < size; i++ {
		p := i - 1
		// 向左找 i 的小于它的第一个值
		for p >= 0 && heights[p] >= heights[i] {
			p = left[p]
		}
		left[i] = p
	}

	// 第 i 个位置的值对应向右找到的小于它的第一个值
	right := make([]int, size)
	right[size-1] = size
	for i := size - 2; i >= 0; i-- {
		p := i + 1
		for p < size && heights[p] >= heights[i] {
			p = right[p]
		}
		right[i] = p
	}

	ans := 0
	for i := 0; i < size; i++ {
		// 注意这里长度要减一
		area := heights[i] * (right[i] - left[i] - 1)
		if area > ans {
			ans = area
		}
	}

	return ans
}
