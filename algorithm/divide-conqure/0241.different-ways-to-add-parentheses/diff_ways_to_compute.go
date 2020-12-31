package problem00241

import "strconv"

var cache map[string][]int = make(map[string][]int)

func diffWaysToCompute(input string) []int {
	result := []int{}
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == '+' || char == '-' || char == '*' {
			leftPart := input[:i]
			rightPart := input[i+1:]
			left, err := cache[leftPart]
			if !err {
				left = diffWaysToCompute(leftPart)
			}
			right, err := cache[rightPart]
			if !err {
				right = diffWaysToCompute(rightPart)
			}
			for _, l := range left {
				for _, r := range right {
					switch char {
					case '+':
						result = append(result, l+r)
						break
					case '-':
						result = append(result, l-r)
						break
					case '*':
						result = append(result, l*r)
						break
					}
				}
			}
		}
	}
	if len(result) == 0 {
		num, _ := strconv.Atoi(input)
		result = append(result, num)
	}
	cache[input] = result
	return result
}
