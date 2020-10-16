package problem0547

import "testing"

func TestFriendCircles(t *testing.T) {
	t.Log(findCircleNum([][]int{
		[]int{1, 0, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 1},
	}))
}
