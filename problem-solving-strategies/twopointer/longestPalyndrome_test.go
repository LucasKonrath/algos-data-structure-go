package twopointer

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Odd length palindrome", "babad", "bab"}, // or "aba"
		{"Single character", "a", "a"},
		{"All same characters", "aaaa", "aaaa"},
		{"No palindrome longer than 1", "abcde", "a"}, // or any single char
		{"Empty string", "", ""},
		{"Palindrome in middle", "abacdfgdcaba", "aba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalyndrome(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("longestPalindrome(%q) = %q, want length %d", tt.input, got, len(tt.want))
			}
		})
	}
}
