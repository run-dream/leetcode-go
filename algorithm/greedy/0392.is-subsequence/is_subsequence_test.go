package problem0392

import "testing"

func TestIsSubsequence(t *testing.T) {
	t.Log(isSubsequence("axc", "ahngdbc") == false)
	t.Log(isSubsequence("aaa", "ahngdbc") == false)
	t.Log(isSubsequence("", "ahngdbc") == true)
}
