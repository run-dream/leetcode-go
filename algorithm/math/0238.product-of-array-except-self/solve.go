package problem0238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	// 求出i左侧的所有乘积
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	// i 右侧的所有乘积
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * right
		right *= nums[i]
	}
	return result
}
