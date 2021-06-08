package offer_merge

// 假设a的实际有效长度为len
func merge(a, b []int, length int) {
	i, j, k := length-1, len(b)-1, len(a)-1
	for k >= 0 && i >= 0 && j >= 0 {
		value := a[i]
		if value > b[j] {
			i--
		} else {
			value = b[j]
			j--
		}
		a[k] = value
		k--
	}
}
