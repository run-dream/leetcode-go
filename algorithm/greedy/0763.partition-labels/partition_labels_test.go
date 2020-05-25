package problem0763

import "testing"

func TestPartitionLabel(t *testing.T) {
	t.Log(SliceEqual(partitionLabels(""), []int{0}) == true)
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
