package helper

import "testing"

func TestToggleBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		pos  uint
		want int
	}{
		{"Toggle 0th bit of 0", 0, 0, 1},
		{"Toggle 1st bit of 0", 0, 1, 2},
		{"Toggle 2nd bit of 0", 0, 2, 4},
		{"Toggle 0th bit of 1", 1, 0, 0},
		{"Toggle 1st bit of 3", 3, 1, 1},
		{"Toggle 2nd bit of 7", 7, 2, 3},
		{"Toggle 3rd bit of 8", 8, 3, 0},
		{"Toggle 4th bit of 8", 8, 4, 24},
		{"Toggle 31st bit of 0", 0, 31, 2147483648},
		{"Toggle 31st bit of 2147483648", 2147483648, 31, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toggleBit(tt.n, tt.pos)
			if got != tt.want {
				t.Errorf("toggleBit(%d, %d) = %d, want %d", tt.n, tt.pos, got, tt.want)
			}
		})
	}
}
