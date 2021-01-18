package problem1025

func divisorGame(N int) bool {
	dp := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			// Bob到达i-j,Alice就能赢
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

func divisorGameGT(N int) bool {
	return N%2 == 0
}
