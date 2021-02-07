package problem0645

func findErrorNums(nums []int) []int {
	n := len(nums)
	visited := make([]bool, n+1)
	repeat := -1
	sum := 0
	for _, val := range nums {
		if visited[val] {
			repeat = val
		} else {
			visited[val] = true
		}
		sum += val
	}
	return []int{repeat, (1+n)*n/2 + repeat - sum}
}
