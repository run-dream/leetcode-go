package problem0067

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	lenA := len(a)
	lenB := len(b)
	extra := 0
	res := []string{}
	for i := lenA - 1; i >= 0; i-- {
		digitA := (int)(a[i] - '0')
		digitB := 0
		indexB := i - (lenA - lenB)
		if indexB >= 0 {
			digitB = (int)(b[indexB] - '0')
		}
		sumI := digitA + digitB + extra
		if sumI > 1 {
			extra = 1
		}else{
			extra = 0
		}
		lastDigit := "0"
		if sumI == 1 || sumI == 3 {
			lastDigit = "1"
		}
		res = append(res, lastDigit)
	}
	if extra == 1 {
		res = append(res, "1")
	}
	result := ""
	for i := len(res) - 1; i >= 0; i-- {
		result += res[i]
	}
	return result
}
