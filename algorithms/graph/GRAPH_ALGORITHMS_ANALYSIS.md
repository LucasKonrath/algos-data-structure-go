# Graph Algorithms - Expert Analysis

## Overview
Graph algorithms form the backbone of many computational problems involving relationships, networks, and connectivity. This comprehensive analysis covers **Breadth-First Search (BFS)**, **Depth-First Search (DFS)**, and **Bellman-Ford Algorithm** with their applications, optimizations, and theoretical foundations.

# Breadth-First Search (BFS)

## Algorithm Overview
BFS explores a graph level by level, visiting all vertices at distance k before visiting any vertex at distance k+1. It uses a **queue** data structure to maintain the order of exploration, guaranteeing the shortest path in unweighted graphs.

## Implementation Analysis

### Core Algorithm
```go
func BFS(graph Graph, start int) []int {
    visited := make(map[int]bool)
    queue := []int{start}
    var result []int
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]  // Dequeue operation
        
        if !visited[node] {
            visited[node] = true
            result = append(result, node)
            
            for _, neighbor := range graph[node] {
                queue = append(queue, neighbor)  // Enqueue neighbors
            }
        }
    }
    return result
}
```

### Algorithm Properties
- **FIFO Exploration**: First discovered vertices are first explored
- **Level-by-Level**: Explores all vertices at distance k before distance k+1
- **Shortest Path**: Guarantees shortest path in unweighted graphs
- **Completeness**: Finds all reachable vertices

## Detailed Execution Trace

### Example Graph
```
Graph representation:
    1 --- 2
    |     |
    3 --- 4 --- 5
          |
          6

Adjacency List:
1: [2, 3]
2: [1, 4]  
3: [1, 4]
4: [2, 3, 5, 6]
5: [4]
6: [4]
```

### BFS Traversal from Node 1
```
Initial: queue = [1], visited = {}, result = []

Step 1: Process node 1
  queue = [], visited = {1}, result = [1]
  Add neighbors: queue = [2, 3]

Step 2: Process node 2  
  queue = [3], visited = {1, 2}, result = [1, 2]
  Add neighbors: queue = [3, 1, 4] (1 already visited)

Step 3: Process node 3
  queue = [1, 4], visited = {1, 2, 3}, result = [1, 2, 3]  
  Add neighbors: queue = [1, 4, 1, 4] (duplicates will be filtered)

Step 4: Process node 1 (already visited, skip)
  queue = [4], visited = {1, 2, 3}, result = [1, 2, 3]

Step 5: Process node 4
  queue = [], visited = {1, 2, 3, 4}, result = [1, 2, 3, 4]
  Add neighbors: queue = [2, 3, 5, 6] (2, 3 already visited)

Step 6: Process remaining nodes 5, 6
  Final result = [1, 2, 3, 4, 5, 6]

Levels:
Level 0: [1]
Level 1: [2, 3]  
Level 2: [4]
Level 3: [5, 6]
```

## Complexity Analysis

### Time Complexity: O(V + E)
```
Vertex Operations: Each vertex visited exactly once = O(V)
Edge Operations: Each edge examined at most twice = O(E)
Queue Operations: O(1) per operation, O(V) total
Total: O(V + E)
```

### Space Complexity: O(V)
```
Visited Set: O(V) to track all vertices
Queue Storage: O(V) in worst case (star graph)
Result Array: O(V) for traversal order
Total: O(V)
```

## BFS Applications and Variants

### Shortest Path in Unweighted Graph
```go
func BFSShortestPath(graph Graph, start, target int) []int {
    if start == target {
        return []int{start}
    }
    
    visited := make(map[int]bool)
    parent := make(map[int]int)
    queue := []int{start}
    visited[start] = true
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                parent[neighbor] = node
                queue = append(queue, neighbor)
                
                if neighbor == target {
                    return reconstructPath(parent, start, target)
                }
            }
        }
    }
    
    return nil // No path found
}

func reconstructPath(parent map[int]int, start, target int) []int {
    path := []int{}
    current := target
    
    for current != start {
        path = append([]int{current}, path...)
        current = parent[current]
    }
    
    return append([]int{start}, path...)
}
```

### Multi-Source BFS
```go
func multiSourceBFS(graph Graph, sources []int) map[int]int {
    distances := make(map[int]int)
    queue := []int{}
    
    // Initialize with all sources
    for _, source := range sources {
        distances[source] = 0
        queue = append(queue, source)
    }
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range graph[node] {
            if _, visited := distances[neighbor]; !visited {
                distances[neighbor] = distances[node] + 1
                queue = append(queue, neighbor)
            }
        }
    }
    
    return distances
}
```

### BFS on Grid (2D Array)
```go
func BFSGrid(grid [][]int, startRow, startCol int) [][]int {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }
    
    queue := [][2]int{{startRow, startCol}}
    visited[startRow][startCol] = true
    
    directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    var result [][2]int
    
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        result = append(result, cell)
        
        for _, dir := range directions {
            newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
            
            if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
               !visited[newRow][newCol] && grid[newRow][newCol] == 1 {
                visited[newRow][newCol] = true
                queue = append(queue, [2]int{newRow, newCol})
            }
        }
    }
    
    return result
}
```

# Depth-First Search (DFS)

## Algorithm Overview
DFS explores a graph by going as deep as possible before backtracking. It uses a **stack** (explicitly or through recursion) to maintain the exploration state, making it ideal for detecting cycles, topological sorting, and connected components.

## Implementation Analysis

### Recursive Implementation
```go
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
```

### Iterative Implementation
```go
func DFSIterative(graph Graph, start int) []int {
    visited := make(map[int]bool)
    stack := []int{start}
    var result []int
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]  // Pop from stack
        
        if !visited[node] {
            visited[node] = true
            result = append(result, node)
            
            // Add neighbors in reverse order for consistent traversal
            neighbors := graph[node]
            for i := len(neighbors) - 1; i >= 0; i-- {
                if !visited[neighbors[i]] {
                    stack = append(stack, neighbors[i])
                }
            }
        }
    }
    
    return result
}
```

## DFS Traversal Patterns

### Pre-order, In-order, Post-order
```go
func DFSWithOrder(graph Graph, start int) ([]int, []int, []int) {
    visited := make(map[int]bool)
    var preOrder, inOrder, postOrder []int
    
    var dfs func(int)
    dfs = func(node int) {
        if visited[node] {
            return
        }
        
        visited[node] = true
        preOrder = append(preOrder, node)  // Pre-order: before recursion
        
        neighbors := graph[node]
        mid := len(neighbors) / 2
        
        // Process left neighbors
        for i := 0; i < mid; i++ {
            if !visited[neighbors[i]] {
                dfs(neighbors[i])
            }
        }
        
        inOrder = append(inOrder, node)  // In-order: middle of recursion
        
        // Process right neighbors
        for i := mid; i < len(neighbors); i++ {
            if !visited[neighbors[i]] {
                dfs(neighbors[i])
            }
        }
        
        postOrder = append(postOrder, node)  // Post-order: after recursion
    }
    
    dfs(start)
    return preOrder, inOrder, postOrder
}
```

## DFS Applications

### Cycle Detection in Directed Graph
```go
func hasCycleDirect(graph Graph) bool {
    WHITE, GRAY, BLACK := 0, 1, 2
    color := make(map[int]int)
    
    var dfs func(int) bool
    dfs = func(node int) bool {
        color[node] = GRAY
        
        for _, neighbor := range graph[node] {
            if color[neighbor] == GRAY {
                return true  // Back edge found, cycle detected
            }
            if color[neighbor] == WHITE && dfs(neighbor) {
                return true
            }
        }
        
        color[node] = BLACK
        return false
    }
    
    // Check all components
    for node := range graph {
        if color[node] == WHITE {
            if dfs(node) {
                return true
            }
        }
    }
    
    return false
}
```

### Topological Sort
```go
func topologicalSort(graph Graph) []int {
    visited := make(map[int]bool)
    var result []int
    
    var dfs func(int)
    dfs = func(node int) {
        visited[node] = true
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
        
        // Add to result in post-order (reverse of finish time)
        result = append([]int{node}, result...)
    }
    
    // Process all nodes
    for node := range graph {
        if !visited[node] {
            dfs(node)
        }
    }
    
    return result
}
```

### Connected Components
```go
func connectedComponents(graph Graph) [][]int {
    visited := make(map[int]bool)
    var components [][]int
    
    var dfs func(int, *[]int)
    dfs = func(node int, component *[]int) {
        visited[node] = true
        *component = append(*component, node)
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor, component)
            }
        }
    }
    
    for node := range graph {
        if !visited[node] {
            var component []int
            dfs(node, &component)
            components = append(components, component)
        }
    }
    
    return components
}
```

# Bellman-Ford Algorithm

## Algorithm Overview
The Bellman-Ford algorithm finds shortest paths from a source vertex to all other vertices in a weighted graph, even with negative edge weights. It can also detect negative cycles, making it more versatile than Dijkstra's algorithm.

## Implementation Analysis

### Core Algorithm
```go
func BellmanFord(n int, edges []Edge, start int) ([]int, bool) {
    // Initialize distances
    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt  // Infinity
    }
    dist[start] = 0
    
    // Relax edges for n-1 iterations
    for i := 0; i < n-1; i++ {
        for _, edge := range edges {
            if dist[edge.from] != math.MaxInt && 
               dist[edge.from]+edge.weight < dist[edge.to] {
                dist[edge.to] = dist[edge.from] + edge.weight
            }
        }
    }
    
    // Check for negative cycles
    for _, edge := range edges {
        if dist[edge.from] != math.MaxInt && 
           dist[edge.from]+edge.weight < dist[edge.to] {
            return nil, false  // Negative cycle detected
        }
    }
    
    return dist, true
}
```

## Algorithm Phases

### Phase 1: Relaxation (n-1 iterations)
```
Relaxation Rule:
if dist[u] + weight(u,v) < dist[v]:
    dist[v] = dist[u] + weight(u,v)

Why n-1 iterations?
- Shortest path has at most n-1 edges
- Each iteration processes paths with one more edge
- After k iterations, all shortest paths with ≤k edges are found
```

### Phase 2: Negative Cycle Detection
```
After n-1 iterations, if any edge can still be relaxed:
→ There exists a negative cycle

Proof:
- If shortest paths exist, they're found in n-1 iterations
- Any further relaxation indicates infinite improvement possible
- This only happens with negative cycles
```

## Detailed Execution Example

### Graph with Negative Edges
```
Vertices: 0, 1, 2, 3
Edges: (0,1,4), (0,2,5), (1,2,-10), (2,3,3), (1,3,5)

Source: 0

Initial: dist = [0, ∞, ∞, ∞]

Iteration 1:
- Relax (0,1): dist[1] = min(∞, 0+4) = 4
- Relax (0,2): dist[2] = min(∞, 0+5) = 5
- Result: dist = [0, 4, 5, ∞]

Iteration 2:
- Relax (1,2): dist[2] = min(5, 4+(-10)) = -6
- Relax (2,3): dist[3] = min(∞, (-6)+3) = -3
- Relax (1,3): dist[3] = min(-3, 4+5) = -3
- Result: dist = [0, 4, -6, -3]

Iteration 3:
- No further improvements possible
- Result: dist = [0, 4, -6, -3]

Negative Cycle Check:
- No edge can be relaxed further
- Return: distances and no negative cycle
```

## Complexity Analysis

### Time Complexity: O(VE)
```
Relaxation Phase: O(V) iterations × O(E) edge checks = O(VE)
Negative Cycle Check: O(E)
Total: O(VE)

Comparison with Dijkstra:
- Dijkstra: O((V + E) log V) with binary heap
- Bellman-Ford: O(VE)
- Trade-off: Bellman-Ford handles negative weights
```

### Space Complexity: O(V)
```
Distance Array: O(V)
Edge List: O(E) (input storage)
Total: O(V + E)
```

## Bellman-Ford Optimizations

### Early Termination
```go
func BellmanFordOptimized(n int, edges []Edge, start int) ([]int, bool) {
    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt
    }
    dist[start] = 0
    
    // Early termination if no changes in iteration
    for i := 0; i < n-1; i++ {
        changed := false
        
        for _, edge := range edges {
            if dist[edge.from] != math.MaxInt && 
               dist[edge.from]+edge.weight < dist[edge.to] {
                dist[edge.to] = dist[edge.from] + edge.weight
                changed = true
            }
        }
        
        if !changed {
            break  // No improvements, algorithm complete
        }
    }
    
    // Negative cycle detection...
    return dist, true
}
```

### SPFA (Shortest Path Faster Algorithm)
```go
func SPFA(graph map[int][]Edge, start int, n int) ([]int, bool) {
    dist := make([]int, n)
    inQueue := make([]bool, n)
    count := make([]int, n)  // Count of times vertex is relaxed
    
    for i := range dist {
        dist[i] = math.MaxInt
    }
    dist[start] = 0
    
    queue := []int{start}
    inQueue[start] = true
    
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        inQueue[u] = false
        
        for _, edge := range graph[u] {
            v := edge.to
            if dist[u] != math.MaxInt && dist[u]+edge.weight < dist[v] {
                dist[v] = dist[u] + edge.weight
                count[v]++
                
                if count[v] >= n {
                    return nil, false  // Negative cycle
                }
                
                if !inQueue[v] {
                    queue = append(queue, v)
                    inQueue[v] = true
                }
            }
        }
    }
    
    return dist, true
}
```

## Advanced Applications

### Currency Arbitrage Detection
```go
func detectArbitrage(rates [][]float64) bool {
    n := len(rates)
    
    // Convert to log space to use addition instead of multiplication
    edges := []Edge{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i != j {
                weight := -math.Log(rates[i][j])  // Negative log
                edges = append(edges, Edge{i, j, int(weight * 1000000)})
            }
        }
    }
    
    // Run Bellman-Ford from any vertex
    _, hasNegativeCycle := BellmanFord(n, edges, 0)
    return !hasNegativeCycle  // Negative cycle means arbitrage opportunity
}
```

### Network Delay Analysis
```go
func networkDelayTime(times [][]int, n int, k int) int {
    edges := []Edge{}
    for _, time := range times {
        edges = append(edges, Edge{time[0]-1, time[1]-1, time[2]})
    }
    
    dist, hasNegativeCycle := BellmanFord(n, edges, k-1)
    if hasNegativeCycle {
        return -1  // Shouldn't happen with delay times
    }
    
    maxTime := 0
    for _, d := range dist {
        if d == math.MaxInt {
            return -1  // Some nodes unreachable
        }
        if d > maxTime {
            maxTime = d
        }
    }
    
    return maxTime
}
```

## Comparison: BFS vs DFS vs Bellman-Ford

### Use Case Matrix
```
Algorithm      | Graph Type    | Edge Weights | Purpose
---------------|---------------|--------------|------------------
BFS            | Any           | Unweighted   | Shortest path, level-order
DFS            | Any           | Any          | Connectivity, cycles, ordering
Bellman-Ford   | Any           | Weighted     | Shortest path with negatives
Dijkstra       | Connected     | Non-negative | Fastest shortest path
Floyd-Warshall | Complete      | Weighted     | All-pairs shortest path
```

### Performance Comparison
```go
func benchmarkGraphAlgorithms(graph Graph, n int) {
    // BFS: O(V + E), good for unweighted shortest paths
    start := time.Now()
    BFS(graph, 0)
    bfsTime := time.Since(start)
    
    // DFS: O(V + E), good for structural analysis
    start = time.Now()
    DFS(graph, 0)
    dfsTime := time.Since(start)
    
    // Bellman-Ford: O(VE), good for negative weights
    edges := graphToEdges(graph)
    start = time.Now()
    BellmanFord(n, edges, 0)
    bfTime := time.Since(start)
    
    fmt.Printf("BFS: %v, DFS: %v, Bellman-Ford: %v\n", 
               bfsTime, dfsTime, bfTime)
}
```

## Error Handling and Edge Cases

### Robust Implementations
```go
func BFSSafe(graph Graph, start int) ([]int, error) {
    if graph == nil {
        return nil, errors.New("graph cannot be nil")
    }
    
    if _, exists := graph[start]; !exists {
        return nil, errors.New("start vertex not in graph")
    }
    
    return BFS(graph, start), nil
}

func BellmanFordSafe(n int, edges []Edge, start int) ([]int, bool, error) {
    if n <= 0 {
        return nil, false, errors.New("number of vertices must be positive")
    }
    
    if start < 0 || start >= n {
        return nil, false, errors.New("start vertex out of range")
    }
    
    // Validate edges
    for _, edge := range edges {
        if edge.from < 0 || edge.from >= n || edge.to < 0 || edge.to >= n {
            return nil, false, errors.New("edge vertices out of range")
        }
    }
    
    return BellmanFord(n, edges, start)
}
```

## Testing Framework
```go
func TestGraphAlgorithms(t *testing.T) {
    // Test graph
    graph := Graph{
        0: {1, 2},
        1: {2, 3},
        2: {3},
        3: {},
    }
    
    // Test BFS
    bfsResult := BFS(graph, 0)
    expectedBFS := []int{0, 1, 2, 3}
    if !reflect.DeepEqual(bfsResult, expectedBFS) {
        t.Errorf("BFS failed: got %v, want %v", bfsResult, expectedBFS)
    }
    
    // Test DFS
    dfsResult := DFS(graph, 0)
    // DFS order depends on implementation, check if all nodes visited
    if len(dfsResult) != 4 {
        t.Errorf("DFS failed: got %d nodes, want 4", len(dfsResult))
    }
    
    // Test Bellman-Ford
    edges := []Edge{
        {0, 1, 1}, {0, 2, 4}, {1, 2, -2}, {2, 3, 3},
    }
    dist, hasNegCycle := BellmanFord(4, edges, 0)
    if hasNegCycle {
        t.Error("Unexpected negative cycle detected")
    }
    expected := []int{0, 1, -1, 2}
    if !reflect.DeepEqual(dist, expected) {
        t.Errorf("Bellman-Ford failed: got %v, want %v", dist, expected)
    }
}
```

These graph algorithms form the foundation for solving complex network problems, from social network analysis to routing protocols, game AI pathfinding, and financial modeling. Understanding their strengths and limitations is crucial for selecting the right approach for specific computational challenges.