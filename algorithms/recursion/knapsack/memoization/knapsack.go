package memoization

import "fmt"

func knapsack(weights []int, values []int, capacity int, n int, memo map[string]int) int {
	if n == 0 || capacity == 0 {
		return 0
	}

	key := fmt.Sprintf("%d,%d", n, capacity)
	if val, exists := memo[key]; exists {
		return val
	}

	if weights[n-1] > capacity {
		memo[key] = knapsack(weights, values, capacity, n-1, memo)
		return memo[key]
	} else {
		includeItem := values[n-1] + knapsack(weights, values, capacity-weights[n-1], n-1, memo)
		excludeItem := knapsack(weights, values, capacity, n-1, memo)
		memo[key] = max(includeItem, excludeItem)
		return memo[key]
	}
}
