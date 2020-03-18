package problem0441

import "testing"

import "math"

func TestArrangeCoins(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(arrangeCoins(i) == int(math.Sqrt(float64(i))))
	}
}
