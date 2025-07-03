package bitmanipulation

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Even positive", 2, true},
		{"Odd positive", 3, false},
		{"Zero", 0, true},
		{"Negative even", -4, true},
		{"Negative odd", -5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isEven(tt.n)
			if got != tt.want {
				t.Errorf("isEven(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
