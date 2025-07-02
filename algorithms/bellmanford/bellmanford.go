package bellmanford

type Edge struct {
	from   int
	to     int
	weight int
}

func BellmanFord(n int, edges []Edge, start int) ([]int, bool) {
	dist := make([]int, n) // Distance from start to each vertex
	for i := range dist {
		dist[i] = int(^uint(0) >> 1) // Initialize distances to infinity
	}
	dist[start] = 0 // Distance to the start vertex is 0

	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			if dist[edge.from] != int(^uint(0)>>1) && dist[edge.from]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[edge.from] + edge.weight
			}
		}
	}

	for _, edge := range edges {
		if dist[edge.from] != int(^uint(0)>>1) && dist[edge.from]+edge.weight < dist[edge.to] {
			return nil, false // Negative cycle detected
		}
	}
	return dist, true // Return distances and indicate no negative cycle
}
