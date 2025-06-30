package doublylinkedlist

import "testing"

func TestDoublyLinkedList_InsertAtEnd(t *testing.T) {
	list := &DoublyLinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)

	current := list.head
	values := []int{}
	for current != nil {
		values = append(values, current.data)
		current = current.next
	}

	expected := []int{10, 20, 30}
	for i, v := range expected {
		if values[i] != v {
			t.Errorf("Expected %d at position %d, got %d", v, i, values[i])
		}
	}
}
