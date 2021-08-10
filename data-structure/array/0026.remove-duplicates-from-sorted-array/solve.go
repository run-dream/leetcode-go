package problem0026

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	slow := 0
	fast := 1
	for fast < size {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
