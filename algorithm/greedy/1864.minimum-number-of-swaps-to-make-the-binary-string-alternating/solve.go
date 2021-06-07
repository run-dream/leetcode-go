package problem1864

func minSwaps(s string) int {
	zeros, ones := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones > zeros+1 || zeros > ones+1 {
		return -1
	}
	if ones > zeros {
		return swap(s, '1')
	} else if zeros > ones {
		return swap(s, '0')
	}
	return min(swap(s, '1'), swap(s, '0'))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func swap(s string, bit byte) int {
	count := 0
	// 通过和标准答案对比，1010101/010101 统计出不正确的位置
	for i := 0; i < len(s); i++ {
		if bit != s[i] {
			count++
		}
		bit = bit ^ 1
	}
	return count / 2
}
