package twopointer

import (
	"testing"
)

func makeList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	var cycleNode *ListNode
	if pos == 0 {
		cycleNode = head
	}
	for i := 1; i < len(vals); i++ {
		n := &ListNode{Val: vals[i]}
		curr.Next = n
		curr = n
		if i == pos {
			cycleNode = n
		}
	}
	if pos >= 0 {
		curr.Next = cycleNode
	}
	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{"No cycle", []int{3, 2, 0, -4}, -1, false},
		{"Cycle at pos 1", []int{3, 2, 0, -4}, 1, true},
		{"Single node, no cycle", []int{1}, -1, false},
		{"Single node, cycle", []int{1}, 0, true},
		{"Empty list", []int{}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := makeList(tt.vals, tt.pos)
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle(%v, pos=%d) = %v, want %v", tt.vals, tt.pos, got, tt.want)
			}
		})
	}
}
