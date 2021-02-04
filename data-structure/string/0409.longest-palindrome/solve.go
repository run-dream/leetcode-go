package problem0409

func longestPalindrome(s string) int {
	m1 := make([]int, 26)
	m2 := make([]int, 26)
	for _, c := range s {
		if c-'a' >= 0 && c-'a' <= 25 {
			m1[c-'a']++
		}
		if c-'A' >= 0 && c-'A' <= 25 {
			m2[c-'A']++
		}
	}
	ans := 0
	odd := 0
	for _, val := range m1 {
		if val%2 == 0 {
			ans += val
		} else {
			ans += val
			odd++
		}
	}
	for _, val := range m2 {
		if val%2 == 0 {
			ans += val
		} else {
			ans += val
			odd++
		}
	}
	if odd > 1 {
		ans -= odd - 1
	}
	return ans
}
