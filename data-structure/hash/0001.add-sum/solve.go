package problem0001

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}
	for i, num := range nums {
		j, found := dict[target-num]
		if found {
			return []int{j, i}
		}
		dict[num] = i
	}
	return []int{-1, -1}
}
