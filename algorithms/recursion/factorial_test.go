package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{10, 3628800},
	}

	for _, test := range tests {
		result := factorial(test.n)
		if result != test.expected {
			t.Errorf("factorial(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}
