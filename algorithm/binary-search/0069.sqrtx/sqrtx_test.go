package problem0069

import "testing"

import "math"

func TestSqrtx(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(mySqrt(i) == int(math.Sqrt(float64(i))))
	}
}

func TestNewtonSqrtx(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(newtonSqrt(i) == int(math.Sqrt(float64(i))))
	}
}
