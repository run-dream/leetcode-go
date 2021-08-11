package problem0189

func rotateNewPlace(nums []int, k int) {
	size := len(nums)
	newNums := make([]int, size)
	for i := 0; i < size; i++ {
		newNums[(i+k)%size] = nums[i]
	}
	copy(nums, newNums)
}

func rotate(nums []int, k int) {
	size := len(nums)
	k = k % size
	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
