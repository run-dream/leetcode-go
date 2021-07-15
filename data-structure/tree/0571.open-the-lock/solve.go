package problem0124

func openLock(deadends []string, target string) int {
	panel := "0000"
	if target == panel {
		return 0
	}
	visited := make(map[string]bool)
	for _, str := range deadends {
		visited[str] = true
	}
	queue := []string{panel}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			visited[cur] = true
			if cur == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := getUp(cur, j)
				if !visited[up] {
					queue = append(queue, up)
					visited[up] = true
				}
				down := getDown(cur, j)
				if !visited[down] {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}

func getUp(str string, index int) string {
	digit := byte(str[index])
	if digit == '9' {
		digit = '0'
	} else {
		digit = byte(str[index] + 1)
	}
	return str[0:index] + string(digit) + str[index+1:]
}

func getDown(str string, index int) string {
	digit := byte(str[index])
	if digit == '0' {
		digit = '9'
	} else {
		digit = byte(str[index] - 1)
	}
	return str[0:index] + string(digit) + str[index+1:]
}
