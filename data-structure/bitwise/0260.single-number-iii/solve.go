package problem0260

func singleNumber(nums []int) []int {
	diff := 0
	// 找到result[0] ^ result[1]的值 因为不同所以 result[1] ^ result[2] != 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	// 找出任意一个不同的位，这里找的是最右侧不同的位
	diff &= -diff
	result := []int{0, 0}
	// 通过这一位把集合分成两部分，一部分是包含diff的集合，另一部分是不包含diff的集合
	for i := 0; i < len(nums); i++ {
		// 这里是不包含的集合，异或的结果就是其中一个元素
		if (nums[i] & diff) == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
