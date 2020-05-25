package problem0763

func partitionLabels(S string) []int {
	lastPosMap := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		lastPosMap[S[i]] = i
	}
	partitions := []int{}
	partition := 0
	for i := 0; i < len(S); i++ {
		lastPos := lastPosMap[S[i]]
		if lastPos > partition {
			partition = lastPos
		}
		if i == partition {
			partitions = append(partitions, i)
		}
	}
	result := make([]int, len(partitions))
	for i := 0; i < len(partitions); i++ {
		start := 0
		if i > 0 {
			start = partitions[i-1] + 1
		}
		result[i] = partitions[i]-start+1
	}
	return result
}
