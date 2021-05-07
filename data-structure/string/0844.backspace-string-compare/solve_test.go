package problem0844

import "testing"

func TestSolve(t *testing.T) {
	t.Log(backspaceCompare("ab#c", "ad#c"))
	t.Log(backspaceCompare("ab##", "c#d#"))
	t.Log(backspaceCompare("a##c", "#a#c"))
	t.Log(backspaceCompare("a#c", "b"))
}
