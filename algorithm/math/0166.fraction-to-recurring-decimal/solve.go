package problem0166

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	result := ""
	// 处理符号位
	digit := ""
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		digit = "-"
	}
	result += digit

	// 整数部分
	numerator = abs(numerator)
	denominator = abs(denominator)

	result += strconv.Itoa(numerator / denominator)
	remainer := numerator % denominator

	if remainer == 0 {
		return result
	}
	// 分数部分
	result += "."
	index := make(map[int]int)
	index[remainer] = len(result)

	for remainer != 0 {
		remainer *= 10
		result += strconv.Itoa(remainer / denominator)
		remainer %= denominator
		// 循环
		if pos, found := index[remainer]; found {
			result = result[0:pos] + "(" + result[pos:] + ")"
			return result
		} else {
			index[remainer] = len(result)
		}
	}
	return result
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
