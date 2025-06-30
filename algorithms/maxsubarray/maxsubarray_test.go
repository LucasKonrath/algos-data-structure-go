package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // [4,-1,2,1]
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-2, -1}, -1},
		{[]int{2, -1, 2, 3, 4, -5}, 10},
		{[]int{-2, 1}, 1},
	}

	for _, test := range tests {
		result := maxSubArray(test.input)
		if result != test.expected {
			t.Errorf("maxSubArray(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}
