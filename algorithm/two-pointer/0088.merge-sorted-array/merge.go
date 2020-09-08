package problem0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < m; i++ {
		nums1[i+m] = nums1[i]
	}
	index1 := m
	index2 := 0
	for index3 := 0; index3 < m+n; index3++ {
		if index2 >= n || (index1 < m+n && nums1[index1] < nums2[index2]) {
			nums1[index3] = nums1[index1]
			index1++
		} else {
			nums1[index3] = nums2[index2]
			index2++
		}
	}
}
