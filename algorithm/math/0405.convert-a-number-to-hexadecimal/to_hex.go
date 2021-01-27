package problem0405

func toHex(num int) string {
	if num < 0 {
		num = 0xFFFFFFFF + num + 1
	}
	if num == 0 {
		return "0"
	}
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	ret := ""
	for num != 0 {
		ret = digits[num%16] + ret
		num = num >> 4
	}
	return ret
}
