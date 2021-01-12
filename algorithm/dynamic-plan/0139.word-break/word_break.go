package problem0139

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			start := i - len(word)
			if start >= 0 && s[start:i] == word {
				dp[i] = dp[i] || dp[start]
			}
		}
	}
	return dp[len(s)]
}
