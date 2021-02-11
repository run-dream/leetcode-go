package problem0442

func solve(nums []int) []int {
	res := []int{}
	for _, val := range nums {
		digit := abs(val) - 1
		if nums[digit] < 0 {
			res = append(res, digit+1)
		} else {
			nums[digit] = -nums[digit]
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
