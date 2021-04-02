package problem1764

func canChoose(groups [][]int, nums []int) bool {
	m := len(groups)
	n := len(nums)
	p := 0
	for i := 0; i < m; i++ {
		group := groups[i]
		matched := false
		for p < n && !matched {
			matched = match(group, nums, p)
			if !matched {
				p++
			} else {
				p += len(group)
			}
		}
		if !matched || p > n {
			return false
		}
	}
	return true
}

func match(group, nums []int, start int) bool {
	for i := 0; i < len(group); i++ {
		pos := i + start
		if pos >= len(nums) {
			return false
		}
		if group[i] != nums[pos] {
			return false
		}
	}
	return true
}
