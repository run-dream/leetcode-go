package problem1441

func buildArray(target []int, n int) []string {
	result := []string{}
	if len(target) == 0 || n == 0 {
		return result
	}
	index := 0
	for i := 1; i <= n; i++ {
		result = append(result, "Push")
		if target[index] != i {
			result = append(result, "Pop")
		} else {
			index++
		}
		if index >= len(target){
			break
		}
	}
	return result
}
