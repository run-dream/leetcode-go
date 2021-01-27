package problem0139

import "testing"

func TestWordBreak(t *testing.T) {
	t.Log(wordBreak("leetcode", []string{"leet", "code"}))
	t.Log(wordBreak("applepenapple", []string{"apple", "pen"}))
	t.Log(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
