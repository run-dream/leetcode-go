package problem0233

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	ans, q, base := 0, n, 1
	for q > 0 {
		digit := q % 10
		q /= 10
		ans += q * base
		if digit == 1 {
			ans += n%base + 1
		} else if digit > 1 {
			ans += base
		}
		base *= 10
	}
	return ans
}
