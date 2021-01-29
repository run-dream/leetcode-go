package problem0739

func dailyTemperatures(T []int) []int {
	stack := make([]int, len(T))
	size := 0
	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for size > 0 && T[i] > T[stack[size-1]] {
			idx := stack[size-1]
			size--
			result[idx] = i - idx
		}
		stack[size] = i
		size++
	}
	return result
}
