package no_repeat_longest_str

import "testing"

func TestSolve(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("abba"))
	t.Log(lengthOfLongestSubstring("pwwkew"))
	t.Log(lengthOfLongestSubstring(""))
}
