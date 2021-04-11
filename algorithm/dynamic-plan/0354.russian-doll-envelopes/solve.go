package problem0354

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	quickSort(envelopes)
	dolls := []int{envelopes[0][1]}
	for i := 1; i < len(envelopes); i++ {
		num := envelopes[i][1]
		if num > dolls[len(dolls)-1] {
			dolls = append(dolls, num)
		} else {
			pos := search(dolls, num)
			dolls[pos] = num
		}
	}
	return len(dolls)
}

func search(dolls []int, target int) int {
	low, high := 0, len(dolls)-1
	for low <= high {
		mid := (low + high) / 2
		if dolls[mid] == target {
			return mid
		} else if dolls[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func compare(a, b []int) bool {
	if a[0] == b[0] {
		return a[1] > b[1]
	}
	return a[0] < b[0]
}

func quickSort(arr [][]int) {
	quickSortRecu(arr, 0, len(arr)-1)
}

func quickSortRecu(arr [][]int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	quickSortRecu(arr, left, pivot-1)
	quickSortRecu(arr, pivot+1, right)
}

func partition(arr [][]int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j <= right-1; j++ {
		if compare(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
