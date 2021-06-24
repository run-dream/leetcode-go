package word_reverse

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	bytes := []byte(s)
	// 去除开头的空白
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != ' ' {
			bytes = bytes[i:]
			break
		}
	}
	// 去除结尾的空白
	for i := len(bytes) - 1; i > 0; i-- {
		if bytes[i] != ' ' {
			bytes = bytes[0 : i+1]
			break
		}
	}
	// 去除重复的空白
	offset := 0
	for i := 1; i < len(bytes); i++ {
		if bytes[i] != ' ' || bytes[i-1] != ' ' {
			bytes[i-offset] = bytes[i]
		} else {
			offset++
		}
	}
	bytes = bytes[0 : len(bytes)-offset]
	// 先整体逆序
	reverseWord(bytes, 0, len(bytes)-1)
	// 再局部逆序
	left := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverseWord(bytes, left, i-1)
			left = i + 1
		}
	}
	reverseWord(bytes, left, len(bytes)-1)
	return string(bytes)
}

func reverseWord(bytes []byte, left, right int) {
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}
