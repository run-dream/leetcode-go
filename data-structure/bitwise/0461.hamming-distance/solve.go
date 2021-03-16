package problem0461

func hammingDistance(x int, y int) int {
	tmp := x ^ y
	res := 0
	for tmp != 0 {
		digit := tmp % 2
		res += digit
		tmp /= 2
	}
	return res
}
