package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{5, 4, 3, 2, 1}, 1, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
		{[]int{7, 7, 7, 7}, 7, 0},
		{[]int{10}, 10, 0},
		{[]int{10}, 5, -1},
	}

	for _, test := range tests {
		result := linearSearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("linearSearch(%v, %d) = %d; want %d", test.arr, test.target, result, test.expected)
		}
	}
}
