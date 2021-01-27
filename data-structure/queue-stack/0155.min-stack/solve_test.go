package problem0155

import "testing"

func TestSolve(t *testing.T) {
	obj := Constructor2()
	obj.Push(0)
	t.Log(obj.Top())
	obj.Push(1)
	t.Log(obj.GetMin())
	obj.Push(0)
	t.Log(obj.GetMin())
	obj.Pop()
	t.Log(obj.GetMin())
}
