package problem0190

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		// 左移一位
		result <<= 1
		// 设置末位为 num的首位
		result |= num & 1
		// num 右移
		num >>= 1
	}
	return result
}
