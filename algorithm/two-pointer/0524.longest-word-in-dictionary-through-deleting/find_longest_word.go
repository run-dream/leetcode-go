package problem0524

func findLongestWord(s string, d []string) string {
	var found string
	for _, item := range d {
		if compare(item, found) && match(s, item) {
			found = item
		}
	}
	return found
}

func match(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	i := 0
	for j := 0; j < len(s2); j++ {
		found := false
		for i < len(s1)-(len(s2)-j)+1 {
			if s1[i] == s2[j] {
				found = true
				i++
				break
			}
			i++
		}
		if !found {
			return false
		}
	}
	return true
}

func compare(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return true
	}
	if len(s1) < len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}
	return false
}
