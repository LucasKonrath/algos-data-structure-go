package memoization

import "testing"

func TestKnapsack(t *testing.T) {
	tests := []struct {
		weights  []int
		values   []int
		capacity int
		n        int
		expected int
	}{
		{[]int{1, 2, 3}, []int{6, 10, 12}, 5, 3, 22},    // Take items 2 and 3
		{[]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5, 4, 7}, // Take items 1 and 2
		{[]int{1, 2, 3}, []int{10, 15, 40}, 6, 3, 65},   // Take all
		{[]int{4, 5, 6}, []int{1, 2, 3}, 3, 3, 0},       // None fit
		{[]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 0, 4, 0}, // Zero capacity
		{[]int{}, []int{}, 10, 0, 0},                    // No items
	}

	for _, test := range tests {
		memo := make(map[string]int)
		result := knapsack(test.weights, test.values, test.capacity, test.n, memo)
		if result != test.expected {
			t.Errorf("knapsack(%v, %v, %d, %d) = %d; want %d", test.weights, test.values, test.capacity, test.n, result, test.expected)
		}
	}
}
