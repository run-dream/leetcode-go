package problem0342

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	//  x&(x-1) 删除最后一位 保证只有一位
	if n&(n-1) != 0 {
		return false
	}
	// '1' 应该在偶数位
	// 去除只是2的倍数的但不是4的倍数的
	return n&0x55555555 != 0
}
