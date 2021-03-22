package problem0318

func maxProduct(words []string) int {
	n := len(words)
	if n == 0 {
		return 0
	}
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			vals[i] |= 1 << int(words[i][j]-'a')
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (vals[i] & vals[j]) == 0 {
				ret = max(ret, len(words[i])*len(words[j]))
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
