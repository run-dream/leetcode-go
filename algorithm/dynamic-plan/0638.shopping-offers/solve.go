package problem0638

import (
	"strconv"
	"strings"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	for i := 0; i < len(price); i++ {
		offer := make([]int, len(price)+1)
		offer[i] = price[i]
		offer[len(price)] = price[i]
		special = append(special, offer)
	}
	memo := make(map[string]int)
	dp(0, special, needs, memo)
}

func toKey(needs []int) string {
	strs := make([]string, len(needs))
	for i := 0; i < len(needs); i++ {
		strs[i] = strconv.Itoa(needs[i])
	}
	return strings.Join(strs, "-")
}

func dp(index int, special [][]int, needs []int, memo map[string]int) int {
	key := toKey(needs)
	if cache, found := memo[key]; found {
		return cache
	}
}
