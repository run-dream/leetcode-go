package problem1331

func arrayRankTransform(arr []int) []int {
	ans := make([]int, len(arr))
	sorted := []int{}
	cached := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		sorted = insert(sorted, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		if cached[arr[i]] != 0 {
			ans[i] = cached[arr[i]]
		} else {
			ans[i] = search(sorted, arr[i]) + 1
			cached[arr[i]] = ans[i]
		}
	}
	return ans
}

func insert(sorted []int, val int) []int {
	low, high := 0, len(sorted)-1
	for low <= high {
		mid := (low + high) / 2
		if sorted[mid] == val {
			return sorted
		}
		if sorted[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return append(sorted[:low], append([]int{val}, sorted[low:]...)...)
}

func search(sorted []int, val int) int {
	low, high := 0, len(sorted)-1
	for low <= high {
		mid := (low + high) / 2
		if sorted[mid] == val {
			return mid
		}
		if sorted[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
