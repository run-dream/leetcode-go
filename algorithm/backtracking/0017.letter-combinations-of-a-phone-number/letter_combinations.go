package problem0017

var digitKeys = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

// Result result
type Result struct {
	result []string
}

func letterCombinations(digits string) []string {
	result := Result{result: []string{}}
	length := len([]byte(digits))
	if length > 0 {
		combinations(digits, "", 0, &result)
	}
	return result.result
}

func combinations(digits string, path string, offset int, result *Result) {
	if offset == len([]byte(digits)) {
		result.result = append(result.result, path)
		return
	}
	for _, char := range digitKeys[digits[offset]] {
		combinations(digits, path+string(char), offset+1, result)
	}
}
