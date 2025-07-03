package twopointer

import (
	"reflect"
	"testing"
)

func TestThreeSums(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{
			name:  "Example 1",
			input: []int{-1, 0, 1, 2, -1, -4},
			want:  [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:  "No triplets",
			input: []int{1, 2, 3, 4, 5},
			want:  [][]int{},
		},
		{
			name:  "All zeros",
			input: []int{0, 0, 0, 0},
			want:  [][]int{{0, 0, 0}},
		},
		{
			name:  "Empty input",
			input: []int{},
			want:  [][]int{},
		},
		{
			name:  "Less than 3 elements",
			input: []int{0, 1},
			want:  [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSums(tt.input)
			if !equalTriplets(got, tt.want) {
				t.Errorf("threeSums(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Helper to compare slices of triplets regardless of order
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, tripletA := range a {
		found := false
		for j, tripletB := range b {
			if !used[j] && reflect.DeepEqual(tripletA, tripletB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
