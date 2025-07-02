package binarysearchtree

type BST struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
	} else {
		insertNode(bst.Root, value)
	}
}

func insertNode(node *Node, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			insertNode(node.Left, value)
		}
	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			insertNode(node.Right, value)
		}
	}
}

func (bst *BST) Search(value int) bool {
	return searchNode(bst.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	} else if value > node.Value {
		return searchNode(node.Right, value)
	}
	return true // value == node.Value
}
