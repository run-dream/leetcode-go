package problem0409

import "testing"

func TestSolve(t *testing.T) {
	t.Log(longestPalindrome("abccccdd"))
	t.Log(longestPalindrome("a"))
	t.Log(longestPalindrome("bb"))
	t.Log(longestPalindrome("ccc"))
}
