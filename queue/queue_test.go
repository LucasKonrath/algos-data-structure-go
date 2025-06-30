package queue

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if len(q.items) != 3 {
		t.Errorf("Expected queue size 3, got %d", len(q.items))
	}

	item, ok := q.Dequeue()
	if !ok || item != 1 {
		t.Errorf("Expected 1, got %v", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != 2 {
		t.Errorf("Expected 2, got %v", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != 3 {
		t.Errorf("Expected 3, got %v", item)
	}

	if len(q.items) != 0 {
		t.Errorf("Expected queue to be empty, got size %d", len(q.items))
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	q := &Queue{}
	item, ok := q.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue on empty queue to return ok=false, got ok=true and item=%v", item)
	}
}

func TestQueue_Front(t *testing.T) {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	item, ok := q.Front()
	if !ok || item != 10 {
		t.Errorf("Expected Front to return 10, got %v", item)
	}
	q.Dequeue()
	item, ok = q.Front()
	if !ok || item != 20 {
		t.Errorf("Expected Front to return 20, got %v", item)
	}
	q.Dequeue()
	_, ok = q.Front()
	if ok {
		t.Error("Expected Front on empty queue to return ok=false")
	}
}
