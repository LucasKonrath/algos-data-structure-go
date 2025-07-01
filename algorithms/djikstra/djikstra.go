package djikstra

import (
	"container/heap"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Graph [][]Edge

type Item struct {
	vertex   int
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{vertex: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.vertex

		if item.distance > dist[u] {
			continue // Skip if this distance is not optimal
		}

		for _, edge := range graph[u] {
			v := edge.to
			weight := edge.weight

			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{vertex: v, distance: dist[v]})
			}
		}
	}
	return dist
}
