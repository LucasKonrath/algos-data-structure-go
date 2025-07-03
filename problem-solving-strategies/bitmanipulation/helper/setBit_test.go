package helper

import "testing"

func TestSetBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		pos  uint
		want int
	}{
		{"Set 0th bit of 0", 0, 0, 1},
		{"Set 1st bit of 0", 0, 1, 2},
		{"Set 2nd bit of 0", 0, 2, 4},
		{"Set 0th bit of 1", 1, 0, 1},
		{"Set 1st bit of 1", 1, 1, 3},
		{"Set 2nd bit of 2", 2, 2, 6},
		{"Set 3rd bit of 8", 8, 3, 8},
		{"Set 4th bit of 8", 8, 4, 24},
		{"Set 31st bit of 0", 0, 31, 2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setBit(tt.n, tt.pos)
			if got != tt.want {
				t.Errorf("setBit(%d, %d) = %d, want %d", tt.n, tt.pos, got, tt.want)
			}
		})
	}
}
