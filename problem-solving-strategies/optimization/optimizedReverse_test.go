package optimization

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"Odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Single element", []int{42}, []int{42}},
		{"Empty slice", []int{}, []int{}},
		{"Negative numbers", []int{-1, -2, -3}, []int{-3, -2, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)
			reverseSlice(inputCopy)
			if !reflect.DeepEqual(inputCopy, tt.want) {
				t.Errorf("reverseSlice(%v) = %v, want %v", tt.input, inputCopy, tt.want)
			}
		})
	}
}
