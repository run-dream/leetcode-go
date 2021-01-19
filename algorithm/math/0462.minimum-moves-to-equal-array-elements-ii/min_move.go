package problem0462

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	ret := 0
	for i <= j {
		ret += nums[j] - nums[i]
		i++
		j--
	}
	return ret
}
