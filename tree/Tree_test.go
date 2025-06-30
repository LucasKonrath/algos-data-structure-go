package tree

import (
	"io"
	"os"
	"strings"
	"testing"
)

func buildTestTree() *Node {
	//      4
	//     / \
	//    2   6
	//   / \ / \
	//  1  3 5  7
	return &Node{
		Value: 4,
		Left: &Node{
			Value: 2,
			Left:  &Node{Value: 1},
			Right: &Node{Value: 3},
		},
		Right: &Node{
			Value: 6,
			Left:  &Node{Value: 5},
			Right: &Node{Value: 7},
		},
	}
}

func captureOutput(f func()) string {
	reader, writer, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = writer

	f()

	writer.Close()
	os.Stdout = old

	var sb strings.Builder
	io.Copy(&sb, reader)
	return sb.String()
}

func TestInOrderTraversal(t *testing.T) {
	tree := buildTestTree()
	output := captureOutput(func() { inOrderTraversal(tree) })
	expected := "1 2 3 4 5 6 7 "
	if output != expected {
		t.Errorf("inOrderTraversal output = %q; want %q", output, expected)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := buildTestTree()
	output := captureOutput(func() { preOrderTraversal(tree) })
	expected := "4 2 1 3 6 5 7 "
	if output != expected {
		t.Errorf("preOrderTraversal output = %q; want %q", output, expected)
	}
}

func TestPostOrderTraversal(t *testing.T) {

	tree := buildTestTree()
	output := captureOutput(func() { postOrderTraversal(tree) })
	expected := "1 3 2 5 7 6 4 "
	if output != expected {
		t.Errorf("postOrderTraversal output = %q; want %q", output, expected)
	}
}
