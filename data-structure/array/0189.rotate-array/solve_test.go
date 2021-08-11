package problem0189

import "testing"

func TestSolve(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateNewPlace(nums, 3)
	t.Log(nums)

	nums = []int{-1, -100, 3, 99}
	rotateNewPlace(nums, 2)
	t.Log(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	t.Log(nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	t.Log(nums)
}
