package bellmanford

import (
	"reflect"
	"testing"
)

func TestBellmanFord_NoNegativeCycle(t *testing.T) {
	edges := []Edge{
		{from: 0, to: 1, weight: 4},
		{from: 0, to: 2, weight: 5},
		{from: 1, to: 2, weight: -2},
		{from: 1, to: 3, weight: 6},
		{from: 2, to: 3, weight: 1},
	}
	dist, ok := BellmanFord(4, edges, 0)
	if !ok {
		t.Fatalf("Expected no negative cycle, but got one")
	}
	expected := []int{0, 4, 2, 3}
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("Expected distances %v, got %v", expected, dist)
	}
}

func TestBellmanFord_NegativeCycle(t *testing.T) {
	edges := []Edge{
		{from: 0, to: 1, weight: 1},
		{from: 1, to: 2, weight: -1},
		{from: 2, to: 0, weight: -1},
	}
	_, ok := BellmanFord(3, edges, 0)
	if ok {
		t.Fatalf("Expected negative cycle, but got none")
	}
}

func TestBellmanFord_SingleNode(t *testing.T) {
	edges := []Edge{}
	dist, ok := BellmanFord(1, edges, 0)
	if !ok {
		t.Fatalf("Expected no negative cycle, but got one")
	}
	expected := []int{0}
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("Expected distances %v, got %v", expected, dist)
	}
}

func TestBellmanFord_UnreachableNodes(t *testing.T) {
	edges := []Edge{
		{from: 0, to: 1, weight: 2},
	}
	dist, ok := BellmanFord(3, edges, 0)
	if !ok {
		t.Fatalf("Expected no negative cycle, but got one")
	}
	inf := int(^uint(0) >> 1)
	expected := []int{0, 2, inf}
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("Expected distances %v, got %v", expected, dist)
	}
}
