package problem0844

func backspaceCompare(s string, t string) bool {
	sResult := getResult(s)
	tResult := getResult(t)
	return sResult == tResult
}

func getResult(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char != '#' {
			result += string(char)
		} else if len(result) > 0 {
			result = result[:len(result)-1]
		}
	}
	return result
}
