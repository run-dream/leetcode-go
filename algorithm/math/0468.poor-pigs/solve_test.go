package problem0468

import "testing"

func TestSolve(t *testing.T) {
	t.Log(poorPigs(1000, 15, 60))
	t.Log(poorPigs(4, 15, 15))
	t.Log(poorPigs(4, 15, 30))
	t.Log(poorPigs(100, 15, 1500))
}
