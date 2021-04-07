package problem0300

func lengthOfLISII(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	list := []int{nums[0]}
	for i := 1; i < n; i++ {
		l, r := 0, len(list)-1
		for l <= r {
			mid := (l + r) / 2
			if nums[i] > list[mid] {
				l++
			} else {
				r--
			}
		}
		if l >= len(list) {
			list = append(list, nums[i])
		} else {
			list[l] = nums[i]
		}
	}
	return len(list)
}
