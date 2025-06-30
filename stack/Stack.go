package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	item := s.items[len(s.items)-1]
	return item, true
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) size() int {
	return len(s.items)
}
