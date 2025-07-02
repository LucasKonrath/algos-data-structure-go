package bfs

type Graph map[int][]int

func BFS(graph Graph, start int) []int {
	visited := make(map[int]bool) // Track visited nodes
	queue := []int{start}         // Initialize queue with the start node
	var result []int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue the first node
		if !visited[node] {
			visited[node] = true // Mark the node as visited
			result = append(result, node)
			for _, neighbor := range graph[node] {
				queue = append(queue, neighbor) // Enqueue all unvisited neighbors
			}
		}
	}
	return result
}
