package fractionalKnapsack

import "sort"

type Item struct {
	value, weight float64
}

func fractionalKnapsack(capacity float64, items []Item) float64 {
	// Sort items by value-to-weight ratio in descending order
	sort.Slice(items, func(i, j int) bool {
		return (items[i].value / items[i].weight) > (items[j].value / items[j].weight)
	})

	totalValue := 0.0
	for _, item := range items {
		if capacity <= 0 {
			break
		}
		if item.weight <= capacity {
			totalValue += item.value
			capacity -= item.weight
		} else {
			totalValue += item.value * (capacity / item.weight)
			capacity = 0
		}
	}
	return totalValue
}
