package problem0020

func isValid(s string) bool {
	stack := make([]byte, len(s))
	size := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[size] = s[i]
			size++
		} else if s[i] == ')' {
			if size == 0 || stack[size-1] != '(' {
				return false
			}
			size--
		} else if s[i] == ']' {
			if size == 0 || stack[size-1] != '[' {
				return false
			}
			size--
		} else if s[i] == '}' {
			if size == 0 || stack[size-1] != '{' {
				return false
			}
			size--
		}
	}
	return size == 0
}
