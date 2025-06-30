package circularlinkedlist

type CircularLinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node // Pointer to the next node
}

func (l *CircularLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		newNode.next = newNode // Point to itself to make it circular
		return
	}
	// Insert at end
	current := l.head
	for current.next != l.head {
		current = current.next
	}
	current.next = newNode
	newNode.next = l.head
}

func (l *CircularLinkedList) DeleteNode(data int) {
	if l.head == nil {
		return
	}
	current := l.head
	prev := l.head
	// Find the node to delete
	for current.data != data {
		if current.next == l.head {
			return // Not found
		}
		prev = current
		current = current.next
	}
	// If only one node
	if current == l.head && current.next == l.head {
		l.head = nil
		return
	}
	// If deleting head
	if current == l.head {
		// Find last node
		tail := l.head
		for tail.next != l.head {
			tail = tail.next
		}
		l.head = l.head.next
		tail.next = l.head
		return
	}
	// Delete non-head node
	prev.next = current.next
}
