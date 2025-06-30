package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
		{[]int{10}, 10, 0},
		{[]int{10}, 5, -1},
		{[]int{1, 3, 5, 7, 9}, 7, 3},
		{[]int{1, 3, 5, 7, 9}, 2, -1},
	}

	for _, test := range tests {
		result := binarySearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("binarySearch(%v, %d) = %d; want %d", test.arr, test.target, result, test.expected)
		}
	}
}
