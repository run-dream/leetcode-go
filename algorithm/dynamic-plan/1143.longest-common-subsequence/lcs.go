package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else if i > 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 0; i < n; i++ {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		} else if i > 0 {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func longestCommonSubsequence2(tex1 string, tex2 string) int {
	text1, text2 := "", ""
	if len(tex1) > len(tex2) {
		text1 = tex1
		text2 = tex2
	} else {
		text2 = tex1
		text1 = tex2
	}

	dpPrev := make([]int, len(text2)+1)
	dpCurr := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dpCurr[j] = dpPrev[j-1] + 1
			} else {
				dpCurr[j] = max(dpCurr[j-1], dpPrev[j])
			}
		}

		dpPrev, dpCurr = dpCurr, dpPrev
	}

	return dpPrev[len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
