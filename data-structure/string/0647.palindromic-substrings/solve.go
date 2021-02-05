package problem0647

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	count := len(s)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = true
			count++
		}
	}
	for size := 3; size <= len(s); size++ {
		for i := size - 1; i < len(s); i++ {
			start := i - size + 1
			if s[start] == s[i] && dp[start+1][i-1] {
				dp[start][i] = true
				count++
			}
		}
	}
	return count
}
