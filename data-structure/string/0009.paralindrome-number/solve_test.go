package problem0009

import "testing"

func TestSolve(t *testing.T) {
	t.Log(isPalindrome(121))
	t.Log(isPalindrome(-1))
	t.Log(isPalindrome(1121))
	t.Log(isPalindrome(1))
	t.Log(isPalindrome(2000))
}
