package problem0476

func findComplement(num int) int {
	mask := 1 << 30
	for (num & mask) == 0 {
		mask >>= 1
	}
	mask = mask<<1 - 1
	return num ^ mask
}
