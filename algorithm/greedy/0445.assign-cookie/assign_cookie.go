package problem0445

func findContentChildren(g []int, s []int) int {
	sort(s)
	sort(g)
	count := 0
	j := 0
	for i := 0; i < len(s) && j < len(g); i++ {
		if s[i] >= g[j] {
			count++
			j++
		}

	}
	return count
}

func sort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i
		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
}
