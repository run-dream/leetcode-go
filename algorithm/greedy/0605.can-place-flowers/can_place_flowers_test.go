package problem0605

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	t.Log((canPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 0) == true))
}
