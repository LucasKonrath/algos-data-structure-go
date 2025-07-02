package dfs

type Graph map[int][]int

func DFS(graph Graph, start int) []int {
	visited := make(map[int]bool)
	var result []int
	var dfsHelper func(int)
	dfsHelper = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		result = append(result, node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}
	dfsHelper(start)
	return result
}
