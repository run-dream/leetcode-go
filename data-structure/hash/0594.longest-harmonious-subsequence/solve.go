package problem0594

func findLHS(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}
	max := 0
	for key, count := range dict {
		if dict[key-1] == 0 {
			continue
		}
		if count+dict[key-1] > max {
			max = count + dict[key-1]
		}
	}
	return max
}
