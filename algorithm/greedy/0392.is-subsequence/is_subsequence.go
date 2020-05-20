package problem0392

func isSubsequence(s string, t string) bool {
	result := true
	i := 0
	for j := 0; j < len(s); j++ {
		found := false
		for ; i < len(t); i++ {
			if s[j] == t[i] {
				found = true
				i++
				break
			}
		}
		if !found {
			result = false
			break
		}
	}
	return result
}
