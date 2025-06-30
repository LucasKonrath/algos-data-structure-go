package deque

import "testing"

func TestDeque_PushPopFrontBack(t *testing.T) {
	d := &Deque{}
	d.PushBack(1)
	d.PushBack(2)
	d.PushFront(0)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	item, ok := d.Front()
	if !ok || item != 0 {
		t.Errorf("Expected Front to return 0, got %v", item)
	}

	item, ok = d.Back()
	if !ok || item != 2 {
		t.Errorf("Expected Back to return 2, got %v", item)
	}

	item, ok = d.PopFront()
	if !ok || item != 0 {
		t.Errorf("Expected PopFront to return 0, got %v", item)
	}

	item, ok = d.PopBack()
	if !ok || item != 2 {
		t.Errorf("Expected PopBack to return 2, got %v", item)
	}

	item, ok = d.PopFront()
	if !ok || item != 1 {
		t.Errorf("Expected PopFront to return 1, got %v", item)
	}

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty")
	}
}

func TestDeque_EmptyOperations(t *testing.T) {
	d := &Deque{}
	if _, ok := d.PopFront(); ok {
		t.Error("Expected PopFront on empty deque to return ok=false")
	}
	if _, ok := d.PopBack(); ok {
		t.Error("Expected PopBack on empty deque to return ok=false")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected Front on empty deque to return ok=false")
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected Back on empty deque to return ok=false")
	}
}
