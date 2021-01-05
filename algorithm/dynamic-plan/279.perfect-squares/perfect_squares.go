package problem0279

func numSquares(n int) int {
	dp := make([]int, n+1)
	nums := getNums(n)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		min := dp[i-1] + 1
		for j := 0; j < len(nums); j++ {
			if i < nums[j] {
				break
			}
			tmp := dp[i-nums[j]] + 1
			if tmp < min {
				min = tmp
			}
		}
		dp[i] = min
	}
	return dp[n]
}

func getNums(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		mul := i * i
		if mul > n {
			break
		}
		result = append(result, mul)
	}
	return result
}
