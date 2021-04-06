package problem0443

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	prev := chars[0]
	prevLen := 1
	result := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] == prev {
			prevLen++
		} else {
			result = adjust(chars, prevLen, result, prev)
			prevLen = 1
			prev = chars[i]
		}
	}
	result = adjust(chars, prevLen, result, prev)
	chars = chars[:result]
	return result
}

func adjust(chars []byte, prevNum, t int, prev byte) int {
	chars[t] = prev
	nums := toBytes(prevNum)
	for i := 0; i < len(nums); i++ {
		chars[t+i+1] = nums[i]
	}
	return t + 1 + len(nums)
}

func toBytes(num int) []byte {
	result := []byte{}
	if num == 1 {
		return result
	}
	for num > 0 {
		result = append(result, byte(num%10+'0'))
		num /= 10
	}
	n := len(result)
	for i := 0; i < n/2; i++ {
		result[i], result[n-i-1] = result[n-i-1], result[i]
	}
	return result
}
