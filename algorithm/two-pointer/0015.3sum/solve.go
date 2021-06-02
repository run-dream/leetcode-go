package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 去重
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				left++
				right--
			}
		}
	}
	return result
}
