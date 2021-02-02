package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)
	for _, char := range s {
		index := int(char - 'a')
		count[index]++
	}
	for _, char := range t {
		index := int(char - 'a')
		count[index]--
		if count[index] < 0 {
			return false
		}
	}
	return true
}
