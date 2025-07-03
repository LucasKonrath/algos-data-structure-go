package twopointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	tortoise, hare := head, head.Next
	for hare != nil && hare.Next != nil {
		if tortoise == hare {
			return true
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return false
}
