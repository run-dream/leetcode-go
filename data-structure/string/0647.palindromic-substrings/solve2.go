package problem0647

func countSubstrings2(s string) int {
	if len(s) <= 0 {
		return 0
	}
	result := 1
	for i := 0; i < len(s)-1; i++ {
		result = expand(i, i, s, result)
		result = expand(i, (i + 1), s, result)
	}
	return result
}

func expand(left int, right int, s string, result int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			result = result + 1
			left = left - 1
			right = right + 1
		} else {
			return result
		}
	}
	return result
}
