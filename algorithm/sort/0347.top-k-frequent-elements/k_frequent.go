package problem0347

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		if val, ok := freqMap[num]; !ok {
			freqMap[num] = 1
		} else {
			freqMap[num] = val + 1
		}
	}
	freqElements := make(map[int][]int)
	for key, val := range freqMap {
		if elems, ok := freqElements[val]; !ok {
			freqElements[val] = []int{key}
		} else {
			freqElements[val] = append(elems, key)
		}
	}
	var result []int
	for i := len(nums); i >= 0 && len(result) < k; i-- {
		if elems, ok := freqElements[i]; ok {
			for _, elem := range elems {
				result = append(result, elem)
			}
		}
	}
	return result
}
