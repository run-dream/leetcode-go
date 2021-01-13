package problem0474

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		count0 := count(str, '0')
		count1 := count(str, '1')
		for j := m; j >= count0; j-- {
			for k := n; k >= count1; k-- {
				if dp[j-count0][k-count1]+1 > dp[j][k] {
					dp[j][k] = dp[j-count0][k-count1] + 1
				}
			}
		}
	}
	return dp[m][n]
}

func count(str string, char byte) int {
	c := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			c++
		}
	}
	return c
}
