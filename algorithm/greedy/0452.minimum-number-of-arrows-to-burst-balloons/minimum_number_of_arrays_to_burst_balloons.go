package problem0452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	arrays := 1
	position := points[0][1]
	for i := 1; i < len(points); i++ {
		// 可以被前一个箭头击破
		if points[i][0] <= position {
			continue
		}
		arrays++
		position = points[i][1]
	}
	return arrays
}
