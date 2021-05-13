package problem0468

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
	return pigs
}
