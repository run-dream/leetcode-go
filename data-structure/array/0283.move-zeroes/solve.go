package problem0283

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j := i + 1
			for j < len(nums) {
				if nums[j] != 0 {
					break
				}
				j++
			}
			if j == len(nums) {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func moveZeroesTwoPointer(nums []int) {
	// 找到第一个为0的位置
	zeroIndex := 0
	for zeroIndex < len(nums) {
		if nums[zeroIndex] == 0 {
			break
		}
		zeroIndex++
	}

	for i := zeroIndex + 1; i < len(nums); i++ {
		if nums[i] != 0 && i > zeroIndex {
			nums[zeroIndex] = nums[i]
			nums[i] = 0
			zeroIndex++
		}
	}
}
