package recursion

import (
	"sort"
	"testing"
)

func sortSolutions(s [][]string) {
	sort.Slice(s, func(i, j int) bool {
		for k := range s[i] {
			if s[i][k] != s[j][k] {
				return s[i][k] < s[j][k]
			}
		}
		return false
	})
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]string
	}{
		{1, [][]string{{"Q"}}},
		{2, [][]string{}},
		{3, [][]string{}},
		{4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
	}

	for _, test := range tests {
		result := solveNQueens(test.n)
		if len(result) != len(test.expected) {
			t.Errorf("solveNQueens(%d) returned %d solutions; want %d", test.n, len(result), len(test.expected))
			continue
		}
		// Sort both for comparison
		sortSolutions(result)
		sortSolutions(test.expected)
		for i := range result {
			for j := range result[i] {
				if result[i][j] != test.expected[i][j] {
					t.Errorf("solveNQueens(%d) solution %d row %d = %q; want %q", test.n, i, j, result[i][j], test.expected[i][j])
				}
			}
		}
	}
}
