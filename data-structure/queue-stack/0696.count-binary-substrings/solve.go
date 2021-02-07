package problem0696

func countBinarySubstrings(s string) int {
	cur := 1
	prev := 0
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
			if prev > 0 && prev >= cur {
				count++
			}
		} else {
			count++
			prev = cur
			cur = 1
		}
	}
	return count
}
