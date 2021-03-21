package problem0693

func hasAlternatingBits(n int) bool {
	// 如果符合条件的话，这里的结果是 11111...
	tmp := n ^ (n >> 1)
	return (tmp & (tmp + 1)) == 0
}
