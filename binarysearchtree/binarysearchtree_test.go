package binarysearchtree

import "testing"

func buildTestBST() *BST {
	bst := &BST{}
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		bst.Insert(v)
	}
	return bst
}

func TestBST_Search(t *testing.T) {
	bst := buildTestBST()

	tests := []struct {
		value    int
		expected bool
	}{
		{8, true},
		{3, true},
		{10, true},
		{1, true},
		{6, true},
		{14, true},
		{4, true},
		{7, true},
		{13, true},
		{0, false},
		{2, false},
		{5, false},
		{9, false},
		{15, false},
	}

	for _, test := range tests {
		result := bst.Search(test.value)
		if result != test.expected {
			t.Errorf("Search(%d) = %v; want %v", test.value, result, test.expected)
		}
	}
}
