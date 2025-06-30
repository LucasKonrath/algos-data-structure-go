package simplylinkedlist

import "fmt"

type Node struct {
	data int
	next *Node // Pointer to the next node
}

type SimplyLinkedList struct {
	head *Node // Pointer to the first node
}

func (l *SimplyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
}

func (l *SimplyLinkedList) Traverse() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
