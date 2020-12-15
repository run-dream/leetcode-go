package problem0093

import "strconv"

func restoreIpAddresses(s string) []string {
	result := []string{}
	for step := 1; step <= 3; step++ {
		result = dfs(s, 0, step, 1, "", result)
	}
	return result
}

func dfs(s string, i, j, k int, ip string, result []string) []string {
	if j > len(s) || k > 3 {
		return result
	}
	newPart := isValidStr(s[i:j])
	if newPart == -1 {
		return result
	}
	newPartStr := strconv.Itoa(newPart)
	if k == 3 {
		part4 := isValidStr(s[j:len(s)])
		if part4 == -1 {
			return result
		}
		ip = ip + "." + newPartStr + "." + strconv.Itoa(part4)
		found := false
		for l := 0; l < len(result); l++ {
			if result[l] == ip {
				found = true
				break
			}
		}
		if !found {
			result = append(result, ip)
		}
		return result
	}
	if k == 1 {
		ip = newPartStr
	} else {
		ip = ip + "." + newPartStr
	}
	for step := 1; step <= 3; step++ {
		result = dfs(s, j, j+step, k+1, ip, result)
	}
	return result
}

func isValidStr(str string) int {
	if len(str) == 0 || len(str) > 4 {
		return -1
	}
	if len(str) > 1 && str[0] == '0' {
		return -1
	}
	var tmp, _ = strconv.Atoi(str)
	if tmp <= 255 && tmp >= 0 {
		return tmp
	}
	return -1
}
