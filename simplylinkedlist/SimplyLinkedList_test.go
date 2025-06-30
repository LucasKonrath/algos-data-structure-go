package simplylinkedlist

import "testing"

func TestSimplyLinkedList_InsertAtBeginning(t *testing.T) {
	list := SimplyLinkedList{}
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)

	current := list.head
	values := []int{}
	for current != nil {
		values = append(values, current.data)
		current = current.next
	}

	expected := []int{30, 20, 10}
	for i, v := range expected {
		if values[i] != v {
			t.Errorf("Expected %d at position %d, got %d", v, i, values[i])
		}
	}
}
