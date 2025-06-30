package doublylinkedlist

import "fmt"

type DoubleNode struct {
	data int
	next *DoubleNode
	prev *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

func (l *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &DoubleNode{data: data}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}
func (l *DoublyLinkedList) ReverseTraverse() {
	current := l.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
}
