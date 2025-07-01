package djikstra

import (
	"math"
	"testing"
)

func TestDijkstra_SimpleGraph(t *testing.T) {
	graph := Graph{
		{{to: 1, weight: 4}, {to: 2, weight: 1}}, // 0
		{{to: 3, weight: 1}},                     // 1
		{{to: 1, weight: 2}, {to: 3, weight: 5}}, // 2
		{},                                       // 3
	}

	dist := Dijkstra(graph, 0)
	expected := []int{0, 3, 1, 4}
	for i, v := range expected {
		if dist[i] != v {
			t.Errorf("vertex %d: expected %d, got %d", i, v, dist[i])
		}
	}
}

func TestDijkstra_DisconnectedGraph(t *testing.T) {
	graph := Graph{
		{{to: 1, weight: 2}}, // 0
		{},                   // 1
		{{to: 3, weight: 1}}, // 2
		{},                   // 3
	}
	dist := Dijkstra(graph, 0)
	if dist[2] != math.MaxInt {
		t.Errorf("vertex 2: expected unreachable (MaxInt), got %d", dist[2])
	}
	if dist[3] != math.MaxInt {
		t.Errorf("vertex 3: expected unreachable (MaxInt), got %d", dist[3])
	}
}

func TestDijkstra_SingleNode(t *testing.T) {
	graph := Graph{{}}
	dist := Dijkstra(graph, 0)
	if dist[0] != 0 {
		t.Errorf("single node: expected 0, got %d", dist[0])
	}
}
