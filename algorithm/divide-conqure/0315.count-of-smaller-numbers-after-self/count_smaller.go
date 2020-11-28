package problem0315

var count []int

func countSmaller(nums []int) []int {
	indexes := make([]int, len(nums))
	for i, _ := range nums {
		indexes[i] = i
	}
	mergeSort(nums, indexes, 0, len(nums)-1)
	return count
}

func mergeSort(nums, indexes []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, indexes, left, mid)
	mergeSort(nums, indexes, mid+1, right)
	merge(nums, indexes, left, right)
}

func merge(nums, indexes []int, left, right int) {
	mid := (left + right) / 2
	leftIndex := left
	rightIndex := mid + 1
	rightCount := 0
	newIndexes := make([]int, right-left+1)
	sortIndex := 0
	for leftIndex <= mid && rightIndex <= right {
		if nums[indexes[leftIndex]] < nums[indexes[rightIndex]] {
			newIndexes[sortIndex] = indexes[leftIndex]
			rightIndex++
			rightCount++
		} else {
			
		}
	}

}
