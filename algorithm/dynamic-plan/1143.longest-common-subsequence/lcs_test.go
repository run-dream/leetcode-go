package problem1143

import "testing"

func TestLCS(t *testing.T) {
	t.Log(longestCommonSubsequence("abcde", "ace") == 3)
	t.Log(longestCommonSubsequence("abc", "abc") == 3)
	t.Log(longestCommonSubsequence("abc", "def") == 0)
	t.Log(longestCommonSubsequence("", "") == 0)
}
