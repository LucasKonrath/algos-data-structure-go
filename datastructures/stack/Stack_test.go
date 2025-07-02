package stack

import "testing"

func TestStack_PushPop(t *testing.T) {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.size())
	}

	item, ok := stack.Pop()
	if !ok || item != 30 {
		t.Errorf("Expected 30, got %v", item)
	}

	item, ok = stack.Pop()
	if !ok || item != 20 {
		t.Errorf("Expected 20, got %v", item)
	}

	item, ok = stack.Pop()
	if !ok || item != 10 {
		t.Errorf("Expected 10, got %v", item)
	}

	if !stack.isEmpty() {
		t.Error("Expected stack to be empty")
	}
}

func TestStack_PopEmpty(t *testing.T) {
	stack := &Stack{}
	item, ok := stack.Pop()
	if ok {
		t.Errorf("Expected Pop on empty stack to return ok=false, got ok=true and item=%v", item)
	}
}
