package no_repeat_longest_str

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	index := make(map[byte]int)
	result := 1
	i := 0
	j := 1
	index[s[0]] = 0
	for j < len(s) {
		// 反向查找 str[j] 最先早出现的位置
		lastPos, found := index[s[j]]
		index[s[j]] = j
		// i 之前的不需要管，因为存在别的字符重复了
		if found && lastPos >= i {
			result = max(result, j-i)
			i = lastPos + 1
		}
		j++
	}
	// 考虑到最后 j = len(s)的情况
	return max(result, j-i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
