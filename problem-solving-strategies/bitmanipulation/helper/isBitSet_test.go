package helper

import "testing"

func TestIsBitSet(t *testing.T) {
	tests := []struct {
		name string
		n    int
		pos  uint
		want bool
	}{
		{"0th bit set in 1", 1, 0, true},
		{"1st bit not set in 1", 1, 1, false},
		{"2nd bit set in 4", 4, 2, true},
		{"2nd bit not set in 2", 2, 2, false},
		{"0th bit not set in 0", 0, 0, false},
		{"3rd bit set in 8", 8, 3, true},
		{"31st bit set in 2147483648", 2147483648, 31, true},
		{"31st bit not set in 0", 0, 31, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBitSet(tt.n, tt.pos)
			if got != tt.want {
				t.Errorf("isBitSet(%d, %d) = %v, want %v", tt.n, tt.pos, got, tt.want)
			}
		})
	}
}
