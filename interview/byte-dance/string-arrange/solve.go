package string_arrage

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	count := map[byte]int{}
	toMatch := 0
	for i := 0; i < len(s1); i++ {
		if count[s1[i]] == 0 {
			toMatch += 1
		}
		count[s1[i]] += 1
	}
	window := map[byte]int{}
	left, right, matched := 0, 0, 0

	for right < len(s2) {
		char := s2[right]
		if count[char] > 0 {
			window[char]++
			if window[char] == count[char] {
				matched++
			}
		}
		right++
		for matched == toMatch && left < right {
			if right-left == len(s1) {
				return true
			}
			char2 := s2[left]
			if count[char2] > 0 {
				if window[char2] == count[char2] {
					matched--
				}
				window[char2]--
			}
			left++
		}
	}
	return false
}
