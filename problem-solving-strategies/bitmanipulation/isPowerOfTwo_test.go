package bitmanipulation

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Zero", 0, false},
		{"Negative number", -2, false},
		{"One", 1, true},
		{"Two", 2, true},
		{"Three", 3, false},
		{"Four", 4, true},
		{"Five", 5, false},
		{"Eight", 8, true},
		{"Sixteen", 16, true},
		{"Large power of two", 1024, true},
		{"Large non-power of two", 1023, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPoweOfTwo(tt.n)
			if got != tt.want {
				t.Errorf("isPoweOfTwo(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
