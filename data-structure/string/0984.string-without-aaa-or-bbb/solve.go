package problem0984

func strWithout3a3b(a int, b int) string {
	isVerse := a < b
	part1 := "aab"
	part2 := "ab"
	part3 := "a"
	if isVerse {
		a, b = b, a
		part1 = "bba"
		part2 = "ba"
		part3 = "b"
	}
	part1Count := a - b
	part2Count := 2*b - a
	part3Count := 0
	for part3Count = 0; part3Count <= 2; part3Count++ {
		ik := part1Count - part3Count
		jk := part2Count + part3Count
		if ik >= 0 && jk >= 0 {
			part1Count = ik
			part2Count = jk
			break
		}
	}
	result := ""
	for i := 0; i < part1Count; i++ {
		result += part1
	}
	for i := 0; i < part2Count; i++ {
		result += part2
	}
	for i := 0; i < part3Count; i++ {
		result += part3
	}
	return result
}
