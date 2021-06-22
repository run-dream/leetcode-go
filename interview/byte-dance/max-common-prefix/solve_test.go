package max_common_prefix

import "testing"

func TestSolve(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
