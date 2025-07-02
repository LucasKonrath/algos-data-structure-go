package Trie

import "testing"

func buildTestTrie() *Trie {
	trie := NewTrie()
	words := []string{"apple", "app", "bat", "bath", "batman", "cat"}
	for _, word := range words {
		trie.Insert(word)
	}
	return trie
}

func TestTrie_Search(t *testing.T) {
	trie := buildTestTrie()

	tests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"app", true},
		{"bat", true},
		{"bath", true},
		{"batman", true},
		{"cat", true},
		{"dog", false},
		{"batm", false},
		{"appl", false},
		{"", false},
	}

	for _, test := range tests {
		result := trie.Search(test.word)
		if result != test.expected {
			t.Errorf("Search(%q) = %v; want %v", test.word, result, test.expected)
		}
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := buildTestTrie()

	tests := []struct {
		prefix   string
		expected bool
	}{
		{"app", true},
		{"bat", true},
		{"bath", true},
		{"batm", true},
		{"cat", true},
		{"c", true},
		{"dog", false},
		{"z", false},
		{"", true},
	}

	for _, test := range tests {
		result := trie.StartsWith(test.prefix)
		if result != test.expected {
			t.Errorf("StartsWith(%q) = %v; want %v", test.prefix, result, test.expected)
		}
	}
}
