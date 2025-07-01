package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},     // 5+5+1
		{[]int{2}, 3, -1},           // not possible
		{[]int{1}, 0, 0},            // zero amount
		{[]int{1}, 2, 2},            // 1+1
		{[]int{5, 10, 25}, 30, 2},   // 25+5
		{[]int{2, 5, 10, 1}, 27, 4}, // 10+10+5+2
	}

	for _, test := range tests {
		result := coinchange(test.coins, test.amount)
		if result != test.expected {
			t.Errorf("coinchange(%v, %d) = %d; want %d", test.coins, test.amount, result, test.expected)
		}
	}
}
