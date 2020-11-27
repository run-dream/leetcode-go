package bitmap

// BitMap 位图
type BitMap struct {
	bytes []byte
	size  int
}

// NewBitMap 新建位图
func NewBitMap(size int) *BitMap {
	return &BitMap{bytes: make([]byte, size/32), size: size}
}

func (m *BitMap) set(k int) {
	if k > m.size {
		return
	}
	byteIndex := uint(k / 8)
	bitIndex := uint(k % 8)
	m.bytes[byteIndex] = byte(uint(m.bytes[byteIndex]) | uint(1<<bitIndex))
}

func (m *BitMap) get(k int) bool {
	if k > m.size {
		return false
	}
	byteIndex := uint(k / 8)
	bitIndex := uint(k % 8)
	return (uint(m.bytes[byteIndex]) & uint(1<<bitIndex)) != 0
}
