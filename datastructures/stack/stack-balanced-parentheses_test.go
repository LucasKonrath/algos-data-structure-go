package stack

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"([]){}", true},
		{"([{}])", true},
		{"((())())", true},
		{"(}", false},
		{"([)]", false},
		{"((()", false},
		{"", true},
		{"[({})]", true},
		{"[({)}]", false},
	}

	for _, test := range tests {
		result := isBalanced(test.input)
		if result != test.expected {
			t.Errorf("isBalanced(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
