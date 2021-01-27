package problem0376

import "testing"

func TestWiggleLength(t *testing.T) {
	t.Log(wiggleMaxLength2([]int{1, 7, 4, 9, 2, 5}))
	t.Log(wiggleMaxLength2([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	t.Log(wiggleMaxLength2([]int{1, 2, 3, 4}))
	t.Log(wiggleMaxLength2([]int{3, 3, 3, 2, 5}))
}
