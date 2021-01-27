package problem0020

import "testing"

func TestSolve(t *testing.T) {
	t.Log(isValid("([]))"))
	t.Log(isValid("[]"))
	t.Log(isValid("{[()]}"))
	t.Log(isValid("{}[]()"))
	t.Log(isValid("({}[])"))
}
