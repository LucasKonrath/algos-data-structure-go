package optimization

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantMin int
		wantMax int
	}{
		{"Empty slice", []int{}, 0, 0},
		{"Single element", []int{5}, 5, 5},
		{"All positive", []int{1, 2, 3, 4, 5}, 1, 5},
		{"All negative", []int{-5, -2, -8, -1}, -8, -1},
		{"Mixed values", []int{-3, 0, 2, 7, -1}, -3, 7},
		{"Duplicates", []int{2, 2, 2, 2}, 2, 2},
		{"Min at end", []int{4, 3, 2, 1, 0}, 0, 4},
		{"Max at start", []int{10, 2, 3, 4, 5}, 2, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := minMax(tt.input)
			if gotMin != tt.wantMin || gotMax != tt.wantMax {
				t.Errorf("minMax(%v) = (%d, %d), want (%d, %d)", tt.input, gotMin, gotMax, tt.wantMin, tt.wantMax)
			}
		})
	}
}
