package bitmanipulation

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Single element", []int{1}, 1},
		{"Two same, one single", []int{2, 2, 1}, 1},
		{"Multiple pairs, one single", []int{4, 1, 2, 1, 2}, 4},
		{"Negative numbers", []int{-1, -1, -2}, -2},
		{"Zero as single", []int{0, 1, 1}, 0},
		{"All pairs, single at end", []int{5, 3, 5, 4, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
