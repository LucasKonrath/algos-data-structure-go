package dfs

import (
	"reflect"
	"testing"
)

func TestDFS_SimpleGraph(t *testing.T) {
	graph := Graph{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {},
	}
	order := DFS(graph, 0)
	expected := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(order, expected) {
		t.Errorf("Expected %v, got %v", expected, order)
	}
}

func TestDFS_DisconnectedGraph(t *testing.T) {
	graph := Graph{
		0: {1},
		1: {},
		2: {3},
		3: {},
	}
	order := DFS(graph, 0)
	expected := []int{0, 1}
	if !reflect.DeepEqual(order, expected) {
		t.Errorf("Expected %v, got %v", expected, order)
	}
}

func TestDFS_SingleNode(t *testing.T) {
	graph := Graph{
		0: {},
	}
	order := DFS(graph, 0)
	expected := []int{0}
	if !reflect.DeepEqual(order, expected) {
		t.Errorf("Expected %v, got %v", expected, order)
	}
}

func TestDFS_EmptyGraph(t *testing.T) {
	graph := Graph{}
	order := DFS(graph, 0)
	expected := []int{0}
	if !reflect.DeepEqual(order, expected) {
		t.Errorf("Expected %v, got %v", expected, order)
	}
}
