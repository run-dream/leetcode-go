package problem0871

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	// 对于每个车站
	for i := 0; i < n && stations[i][0] < target; i++ {
		for j := i + 1; j > 0; j-- {
			// 如果 j-1 次可以经过stations[i]，那么第j次就可以经过dp[j-1] + stations[i].fuel
			if dp[j-1] >= stations[i][0] {
				dp[j] = max(dp[j], dp[j-1]+stations[i][1])
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
