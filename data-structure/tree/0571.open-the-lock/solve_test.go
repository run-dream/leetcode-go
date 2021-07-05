package problem0124

import "testing"

func TestSolve(t *testing.T) {
	t.Log(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
