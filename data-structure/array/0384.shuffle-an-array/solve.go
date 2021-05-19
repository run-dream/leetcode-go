package problem0384

import "math/rand"

type Solution struct {
	// 原数组
	nums []int
	size int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums,
		len(nums),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	result := make([]int, this.size)
	copy(result, this.nums)
	for i := this.size - 1; i >= 0; i-- {
		j := rand.Intn(this.size)
		result[i], result[j] = result[j], result[i]
	}
	return result
}
