package problem0605

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		unplaced := flowerbed[i] == 0
		leftUnplaced := i == 0 || (i > 0 && flowerbed[i-1] == 0)
		rightUnplaced := (i == len(flowerbed)-1) || (i < (len(flowerbed)-1) && flowerbed[i+1] == 0)
		if unplaced && leftUnplaced && rightUnplaced {
			flowerbed[i] = 1
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}
