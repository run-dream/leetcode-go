package problem0169

func majorityElement(nums []int) int {
	cnt := 0
	majority := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			majority = nums[i]
			cnt++
		} else if majority == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return majority
}
