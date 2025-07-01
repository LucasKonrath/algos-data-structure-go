package coinchange

import "sort"

func coinchange(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Sort coins in descending order
	count := 0
	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			count++
		}
	}
	if amount == 0 {
		return count
	}
	return -1
}
