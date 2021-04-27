package problem0717

import "testing"

func TestSolve(t *testing.T) {
	t.Log(isOneBitCharacter([]int{1, 0, 0}))
	t.Log(isOneBitCharacter([]int{1, 1, 1, 0}))
}
