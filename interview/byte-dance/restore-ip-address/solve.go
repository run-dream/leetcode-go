package restore_ip_address

import "strconv"

func restoreIpAddresses(s string) []string {
	result := []string{}
	if len(s) == 0 {
		return result
	}
	for size := 1; size <= 3; size++ {
		restoreIpAddress(&result, s, "", 1, 0, size)
	}
	return result
}

func restoreIpAddress(result *[]string, s, path string, step, begin, length int) {
	if step > 3 || begin+length > len(s) {
		return
	}
	newPart := s[begin : begin+length]
	if !isValid(newPart) {
		return
	}
	if step == 1 {
		path = newPart
	} else {
		path = path + "." + newPart
	}

	if step == 3 {
		part4 := s[begin+length:]
		if isValid(part4) {
			*result = append(*result, path+"."+part4)
		}
		return
	}

	for i := 1; i <= 3; i++ {
		restoreIpAddress(result, s, path, step+1, begin+length, i)
	}
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if s[0] == '0' {
		return false
	}
	val, _ := strconv.Atoi(s)
	return val >= 0 && val <= 255
}
