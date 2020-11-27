package bitmap

import "testing"

func TestBitMap(t *testing.T) {
	bitmap := NewBitMap(2 << 32)
	bitmap.set(1023)
	t.Log(bitmap.get(111024))
	t.Log(bitmap.get(1023))
}
