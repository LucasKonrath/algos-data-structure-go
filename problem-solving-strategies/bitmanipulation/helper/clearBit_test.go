package helper

import "testing"

func TestClearBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		pos  uint
		want int
	}{
		{"Clear 0th bit of 1", 1, 0, 0},
		{"Clear 1st bit of 3", 3, 1, 1},
		{"Clear 2nd bit of 7", 7, 2, 3},
		{"Clear 0th bit of 0", 0, 0, 0},
		{"Clear 3rd bit of 8", 8, 3, 0},
		{"Clear 4th bit of 24", 24, 4, 8},
		{"Clear 31st bit of 2147483648", 2147483648, 31, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clearBit(tt.n, tt.pos)
			if got != tt.want {
				t.Errorf("clearBit(%d, %d) = %d, want %d", tt.n, tt.pos, got, tt.want)
			}
		})
	}
}
