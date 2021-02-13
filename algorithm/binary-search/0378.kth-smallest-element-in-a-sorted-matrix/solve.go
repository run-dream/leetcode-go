package problem0378

func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	low := matrix[0][0]
	high := matrix[m-1][n-1]
	for low < high {
		mid := low + (high-low)/2
		count := 0
		j := n - 1
		for i := 0; i < m; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += (j + 1)
		}
		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
