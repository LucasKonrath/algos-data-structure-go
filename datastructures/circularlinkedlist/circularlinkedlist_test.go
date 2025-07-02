package circularlinkedlist

import (
	"testing"
)

func TestCircularLinkedList_DeleteNode(t *testing.T) {
	list := &CircularLinkedList{}
	// Insert nodes
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)

	// Delete head node
	list.DeleteNode(1)
	if list.head.data != 2 {
		t.Errorf("Expected head to be 2, got %d", list.head.data)
	}

	// Delete middle node
	list.DeleteNode(3)
	current := list.head
	found := false
	for i := 0; i < 3; i++ {
		if current.data == 3 {
			found = true
			break
		}
		current = current.next
	}
	if found {
		t.Error("Node with data 3 should have been deleted")
	}

	// Delete last node
	list.DeleteNode(4)
	current = list.head
	for i := 0; i < 2; i++ {
		if current.data == 4 {
			t.Error("Node with data 4 should have been deleted")
		}
		current = current.next
	}

	// Delete non-existent node
	list.DeleteNode(99) // Should not panic or change the list
	if list.head == nil {
		t.Error("List head should not be nil after deleting non-existent node")
	}
	if list.head.data != 2 || list.head.next != list.head {
		t.Error("List structure changed after deleting non-existent node")
	}
}
