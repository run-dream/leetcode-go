package problem1680

func concatenatedBinary(n int) int {
	var result int64 = 0
	var bits int64 = 1
	var M int64 = 1E9 + 7
	var i int64 = 1
	for int(i) <= n {
		result = ((result << bits) + i) % M
		if i == ((1 << bits) - 1) {
			bits++
		}
		i++
	}
	return int(result)
}
