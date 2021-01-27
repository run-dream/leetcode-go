package problem0415

func addStrings(num1 string, num2 string) string {
	carry := 0
	result := ""
	i := len(num1) - 1
	j := len(num2) - 1
	for i >= 0 || j >= 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		result = string((byte)('0'+carry%10)) + result
		carry /= 10
	}
	if carry != 0 {
		result = string((byte)('0'+carry)) + result
	}
	return result
}
