package quick_sort

func quickSort(arr []int) {
	quickSortRecu(arr, 0, len(arr)-1)
}

func quickSortRecu(arr []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	quickSortRecu(arr, left, pivot-1)
	quickSortRecu(arr, left+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
