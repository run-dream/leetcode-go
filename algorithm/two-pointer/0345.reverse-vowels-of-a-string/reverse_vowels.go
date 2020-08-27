package problem0345

func reverseVowels(s string) string {
	arr := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isVowel(arr[i]) {
			i++
		} else if !isVowel(arr[j]) {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return string(arr)
}

func isVowel(char byte) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == char {
			return true
		}
	}
	return false
}
