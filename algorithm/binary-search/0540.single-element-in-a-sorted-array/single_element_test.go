package problem0540

import "testing"

func TestSingleElement(t *testing.T) {
	t.Log(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}) == 2)
	t.Log(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}) == 10)
}
