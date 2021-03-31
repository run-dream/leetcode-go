package problem0507

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	i := 2
	for ; i*i < num; i++ {
		if num%i == 0 {
			sum += num/i + i
		}
	}
	if i*i == num {
		sum += i
	}
	return sum == num
}
