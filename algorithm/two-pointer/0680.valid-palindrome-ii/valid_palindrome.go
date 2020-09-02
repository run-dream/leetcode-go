package problem0680

func validPalindrome(s string) bool {
	return isPalindrome(s, 0, len(s)-1, false)
}

func isPalindrome(s string, left, right int, isDeleted bool) bool {
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else if isDeleted {
			return false
		} else {
			return isPalindrome(s, left+1, right, true) || isPalindrome(s, left, right-1, true)
		}
	}
	return true
}
