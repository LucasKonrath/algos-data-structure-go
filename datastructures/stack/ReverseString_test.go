package stack

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
		{"GoLang", "gnaLoG"},
		{"12345", "54321"},
		{"!@#", "#@!"},
		{"你好", "好你"},
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
