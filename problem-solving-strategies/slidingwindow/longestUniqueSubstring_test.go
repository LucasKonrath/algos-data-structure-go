package slidingwindow

import "testing"

func TestLongestUniquesSubstring_Basic(t *testing.T) {
	input := "abcabcbb"
	expected := "abc"
	result := longestUniquesSubstring(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestLongestUniquesSubstring_AllUnique(t *testing.T) {
	input := "abcdef"
	expected := "abcdef"
	result := longestUniquesSubstring(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestLongestUniquesSubstring_AllSame(t *testing.T) {
	input := "aaaaaa"
	expected := "a"
	result := longestUniquesSubstring(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestLongestUniquesSubstring_Empty(t *testing.T) {
	input := ""
	expected := ""
	result := longestUniquesSubstring(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestLongestUniquesSubstring_MiddleLongest(t *testing.T) {
	input := "abcbdeaf"
	expected := "cbdeaf"
	result := longestUniquesSubstring(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
