package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubblesort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 3, 2, 1, 1}, []int{1, 1, 2, 2, 3}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
	}

	for _, test := range tests {
		arr := make([]int, len(test.input))
		copy(arr, test.input)
		bubblesort(arr)
		if !reflect.DeepEqual(arr, test.expected) {
			t.Errorf("bubblesort(%v) = %v; want %v", test.input, arr, test.expected)
		}
	}
}
