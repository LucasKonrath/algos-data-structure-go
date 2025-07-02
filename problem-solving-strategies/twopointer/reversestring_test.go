package twopointer

import "testing"

func TestReverseString_Normal(t *testing.T) {
	input := "hello"
	expected := "olleh"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_Empty(t *testing.T) {
	input := ""
	expected := ""
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_SingleChar(t *testing.T) {
	input := "a"
	expected := "a"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestReverseString_Palindrome(t *testing.T) {
	input := "madam"
	expected := "madam"
	result := reverseString(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
