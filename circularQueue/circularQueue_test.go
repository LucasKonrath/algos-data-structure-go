package circularQueue

import "testing"

func TestCircularQueue_EnqueueDequeue(t *testing.T) {
	cq := NewCircularQueue(3)
	if !cq.Enqueue(1) {
		t.Error("Expected Enqueue(1) to succeed")
	}
	if !cq.Enqueue(2) {
		t.Error("Expected Enqueue(2) to succeed")
	}
	if !cq.Enqueue(3) {
		t.Error("Expected Enqueue(3) to succeed")
	}
	if cq.Enqueue(4) {
		t.Error("Expected Enqueue(4) to fail (queue should be full)")
	}

	if !cq.IsFull() {
		t.Error("Expected queue to be full")
	}

	item, ok := cq.Dequeue()
	if !ok || item != 1 {
		t.Errorf("Expected Dequeue to return 1, got %v", item)
	}
	if cq.IsFull() {
		t.Error("Expected queue to not be full after Dequeue")
	}

	if !cq.Enqueue(4) {
		t.Error("Expected Enqueue(4) to succeed after Dequeue")
	}

	item, ok = cq.Front()
	if !ok || item != 2 {
		t.Errorf("Expected Front to return 2, got %v", item)
	}

	item, ok = cq.Dequeue()
	if !ok || item != 2 {
		t.Errorf("Expected Dequeue to return 2, got %v", item)
	}
	item, ok = cq.Dequeue()
	if !ok || item != 3 {
		t.Errorf("Expected Dequeue to return 3, got %v", item)
	}
	item, ok = cq.Dequeue()
	if !ok || item != 4 {
		t.Errorf("Expected Dequeue to return 4, got %v", item)
	}
	if !cq.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}

func TestCircularQueue_EmptyDequeueFront(t *testing.T) {
	cq := NewCircularQueue(2)
	item, ok := cq.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue on empty queue to return ok=false, got ok=true and item=%v", item)
	}
	item, ok = cq.Front()
	if ok {
		t.Errorf("Expected Front on empty queue to return ok=false, got ok=true and item=%v", item)
	}
}
