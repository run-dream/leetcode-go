package problem0232

import "testing"

func TestSolve(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	t.Log(obj.Pop())
	obj.Push(2)
	t.Log(obj.Peek())
	t.Log(obj.Empty())
}
