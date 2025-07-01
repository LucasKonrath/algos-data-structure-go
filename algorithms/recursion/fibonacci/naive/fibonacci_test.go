package naive

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := fibonacci(test.n)
		if result != test.expected {
			t.Errorf("fibonacci(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}
