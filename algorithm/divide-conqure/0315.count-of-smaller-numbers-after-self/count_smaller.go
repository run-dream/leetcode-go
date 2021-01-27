package problem0315

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	indexes := make([]int, len(nums))
	for i := range nums {
		indexes[i] = i
	}
	mergeSort(nums, &count, &indexes, 0, len(nums)-1)
	return count
}

func mergeSort(nums []int, count, indexes *[]int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, count, indexes, left, mid)    // 排序左侧
	mergeSort(nums, count, indexes, mid+1, right) // 排序右侧
	merge(nums, count, indexes, left, right)      // 合并
}

func merge(nums []int, count, indexes *[]int, left, right int) { // 右侧有几个数比它小
	merged := make([]int, right-left+1) // 下标转化数组
	mid := (left + right) / 2
	leftCursor := left     //  左侧游标
	rightCursor := mid + 1 // 右侧游标
	mergedCursor := 0
	rightCount := 0
	for leftCursor <= mid || rightCursor <= right { // 能合并
		if leftCursor > mid {
			// 左侧的数组已经处理完了，直接将右侧数组加进来
			rightCount++
			merged[mergedCursor] = (*indexes)[rightCursor]
			rightCursor++
		} else if rightCursor > right {
			//  右侧数组处理完了
			(*count)[(*indexes)[leftCursor]] = (*count)[(*indexes)[leftCursor]] + rightCount
			merged[mergedCursor] = (*indexes)[leftCursor]
			leftCursor++
		} else if nums[(*indexes)[leftCursor]] <= nums[(*indexes)[rightCursor]] {
			(*count)[(*indexes)[leftCursor]] += rightCount
			merged[mergedCursor] = (*indexes)[leftCursor]
			leftCursor++
		} else {
			rightCount++
			merged[mergedCursor] = (*indexes)[rightCursor]
			rightCursor++
		}
		mergedCursor++
	}

	for i := 0; i < len(merged); i++ {
		(*indexes)[left+i] = merged[i]
	}

}
