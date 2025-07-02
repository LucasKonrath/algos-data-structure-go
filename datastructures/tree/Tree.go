package tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left)
	fmt.Print(node.Value, " ")
	inOrderTraversal(node.Right)
}

func preOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}

func postOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Print(node.Value, " ")
}
