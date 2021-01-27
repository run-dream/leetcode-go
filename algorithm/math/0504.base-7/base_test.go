package problem0504

import "testing"

func TestBase(t *testing.T) {
	t.Log(convertToBase7(1))
	t.Log(convertToBase7(100))
	t.Log(convertToBase7(-100))
}
