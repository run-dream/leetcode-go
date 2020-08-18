package problem0406

func reconstructQueue(people [][]int) [][]int {
	sort(people)
	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}
	queued := make([][]int, len(people))
	for _, v := range people {
		queued[indexes[v[1]]] = v
		indexes = append(indexes[:v[1]], indexes[v[1]+1:]...)
	}
	return queued
}

func sort(people [][]int) {
	length := len(people)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i
		for j := length - 1; j > i; j-- {
			if people[j][1] < people[min][1] || (people[j][1] == people[min][1] && people[j][0] > people[min][0]) {
				min = j
			}
		}
		people[i], people[min] = people[min], people[i]
	}
}
