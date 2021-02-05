package problem0009

func isPalindrome(x int) bool {
	if x < 0 || (x > 9 && x%10 == 0) {
		return false
	}
	orig := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev == orig
}
