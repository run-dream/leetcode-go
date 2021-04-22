package problem0400

func findNthDigit(n int) int {
	l := 1
	i := 1
	iDigit := 1
	thro := iDigit * 10
	for l < n {
		i++
		if thro <= i {
			iDigit++
			thro *= 10
		}
		l += iDigit
	}
	pos := l - n + 1
	res := i
	for j := 0; j < pos; j++ {
		res = i % 10
		i /= 10
	}
	return res
}
