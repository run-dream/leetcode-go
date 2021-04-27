package problem0717

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	for i := 0; i < n; {
		bit := bits[i]
		if i == n-1 {
			return true
		}
		if bit == 1 {
			i += 2
		} else {
			i++
		}
	}
	return false
}
