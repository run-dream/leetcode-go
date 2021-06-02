package problem0057

func insert(intervals [][]int, newInterval []int) [][]int {
	startIndex, mergeLen, prevIndex := -1, 0, -1
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if isIntersect(interval, newInterval) {
			if startIndex == -1 {
				startIndex = i
			}
			mergeLen++
			newInterval = merge(interval, newInterval)
		}
		if interval[1] < newInterval[0] {
			prevIndex = i
		}
	}
	if mergeLen == 0 {
		return append(intervals[0:prevIndex+1], append([][]int{newInterval}, intervals[prevIndex+1:]...)...)
	}
	return append(intervals[0:startIndex], append([][]int{newInterval}, intervals[startIndex+mergeLen:]...)...)
}

func isIntersect(interval, newInterval []int) bool {
	// 如果两个区间不相交，那么最大的开始端一定大于最小的结束端
	maxBegin := interval[0]
	if newInterval[0] > maxBegin {
		maxBegin = newInterval[0]
	}
	minEnd := interval[1]
	if newInterval[1] < minEnd {
		minEnd = newInterval[1]
	}
	return maxBegin <= minEnd
}

func merge(interval, newInterval []int) []int {
	minBegin := interval[0]
	if newInterval[0] < minBegin {
		minBegin = newInterval[0]
	}
	maxEnd := interval[1]
	if newInterval[1] > maxEnd {
		maxEnd = newInterval[1]
	}
	return []int{minBegin, maxEnd}
}
