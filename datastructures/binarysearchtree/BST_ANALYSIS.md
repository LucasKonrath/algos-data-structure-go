# Binary Search Tree (BST) - Expert Analysis

## Data Structure Overview
A **Binary Search Tree** is a hierarchical data structure where each node has at most two children, and the **BST property** is maintained: for any node, all values in the left subtree are smaller, and all values in the right subtree are larger.

## BST Property (Invariant)
For every node n in the tree:
- **Left Subtree**: ∀ node ∈ left_subtree(n), node.value < n.value
- **Right Subtree**: ∀ node ∈ right_subtree(n), node.value > n.value
- **No Duplicates**: This implementation ignores duplicate insertions

## Implementation Analysis

### Node Structure
```go
type Node struct {
    Value int   // Stored data
    Left  *Node // Left child pointer
    Right *Node // Right child pointer
}

type BST struct {
    Root *Node  // Tree root reference
}
```

### Memory Layout
```
Node Memory Footprint:
├── Value: 8 bytes (int64 on 64-bit systems)
├── Left:  8 bytes (pointer)
└── Right: 8 bytes (pointer)
Total: 24 bytes per node + heap allocation overhead
```

## Operation Analysis

### Insertion Algorithm
```go
func (bst *BST) Insert(value int) {
    if bst.Root == nil {
        bst.Root = &Node{Value: value}
    } else {
        insertNode(bst.Root, value)
    }
}
```

#### Insertion Trace Example
```
Insert sequence: [5, 3, 7, 2, 4, 6, 8]

Step 1: Insert 5
    5
   
Step 2: Insert 3
    5
   /
  3

Step 3: Insert 7  
    5
   / \
  3   7

Step 4: Insert 2
    5
   / \
  3   7
 /
2

Final Tree:
      5
    /   \
   3     7
  / \   / \
 2   4 6   8
```

### Search Algorithm
```go
func searchNode(node *Node, value int) bool {
    if node == nil {
        return false
    }
    if value < node.Value {
        return searchNode(node.Left, value)
    } else if value > node.Value {
        return searchNode(node.Right, value)
    }
    return true
}
```

#### Search Path Analysis
For a balanced BST with height h:
- **Comparisons**: O(h) = O(log n) average case
- **Memory Access**: O(h) cache misses for large trees
- **Branch Predictions**: Generally good due to structured traversal

## Complexity Analysis

### Time Complexity
| Operation | Best Case | Average Case | Worst Case |
|-----------|-----------|--------------|------------|
| Search    | O(1)      | O(log n)     | O(n)       |
| Insert    | O(1)      | O(log n)     | O(n)       |
| Delete*   | O(1)      | O(log n)     | O(n)       |

*Delete operation not implemented in current code

### Space Complexity
- **Storage**: O(n) for n nodes
- **Recursion Stack**: O(h) where h is tree height
- **Worst Case Stack**: O(n) for skewed tree

## Tree Shape Analysis

### Balanced Tree Properties
```
Height: h = ⌊log₂(n)⌋
Nodes at level i: 2^i
Maximum nodes: 2^(h+1) - 1
Minimum nodes for height h: h + 1
```

### Degenerate Cases

#### Right-Skewed Tree (Sorted Input)
```
Insert [1, 2, 3, 4, 5]:
1
 \
  2
   \
    3
     \
      4
       \
        5
Height = n-1, Operations = O(n)
```

#### Left-Skewed Tree (Reverse Sorted)
```
Insert [5, 4, 3, 2, 1]:
    5
   /
  4
 /
3
/
2
/
1
Height = n-1, Operations = O(n)
```

## Tree Traversal Algorithms

### In-Order Traversal (Sorted Output)
```go
func inOrderTraversal(node *Node) []int {
    if node == nil {
        return []int{}
    }
    
    result := []int{}
    result = append(result, inOrderTraversal(node.Left)...)
    result = append(result, node.Value)
    result = append(result, inOrderTraversal(node.Right)...)
    return result
}
```

### Traversal Types
- **Pre-Order**: Root → Left → Right (Tree structure copying)
- **In-Order**: Left → Root → Right (Sorted sequence)  
- **Post-Order**: Left → Right → Root (Tree deletion)
- **Level-Order**: BFS traversal (Tree visualization)

## Mathematical Properties

### Height Distribution
For random insertion sequence of n distinct keys:
- **Expected Height**: E[h] ≈ 2.99 log n
- **Variance**: Var[h] ≈ 1.38
- **Probability of Balance**: P(h ≤ c log n) → 1 as n → ∞

### Path Length Analysis
```
Internal Path Length (IPL):
IPL = Σ(depth(node) × 1) for all internal nodes

External Path Length (EPL):  
EPL = Σ(depth(leaf) × 1) for all leaves

Relationship: EPL = IPL + 2n
```

## Performance Optimization Strategies

### Tree Balancing
```go
// Self-balancing alternatives:
// - AVL Trees: Height difference ≤ 1
// - Red-Black Trees: Relaxed balancing with color properties
// - Splay Trees: Recently accessed nodes move to root
```

### Cache Optimization
```go
// Memory layout optimization for better cache performance
type CacheOptimizedNode struct {
    Value int
    Left, Right uint32 // Array indices instead of pointers
}

type CompactBST struct {
    Nodes []CacheOptimizedNode
    Root  uint32
}
```

### Bulk Operations
```go
func BuildBalancedBST(sortedArray []int) *BST {
    // Build balanced BST from sorted array in O(n) time
    return buildBSTFromSorted(sortedArray, 0, len(sortedArray)-1)
}

func buildBSTFromSorted(arr []int, start, end int) *Node {
    if start > end {
        return nil
    }
    
    mid := (start + end) / 2
    node := &Node{Value: arr[mid]}
    
    node.Left = buildBSTFromSorted(arr, start, mid-1)
    node.Right = buildBSTFromSorted(arr, mid+1, end)
    
    return node
}
```

## Advanced BST Variants

### Threaded Binary Trees
```go
type ThreadedNode struct {
    Value       int
    Left, Right *ThreadedNode
    LeftThread  bool // true if Left points to predecessor
    RightThread bool // true if Right points to successor
}
```

### Order Statistics Trees
```go
type OrderStatNode struct {
    Value int
    Left, Right *OrderStatNode
    Size  int // Number of nodes in subtree
}

// Find k-th smallest element in O(log n)
func (ost *OrderStatTree) Select(k int) int
// Find rank of element in O(log n)  
func (ost *OrderStatTree) Rank(value int) int
```

## Implementation Improvements

### Missing Operations
```go
// Deletion with three cases:
// 1. Leaf node: Simply remove
// 2. One child: Replace with child
// 3. Two children: Replace with inorder successor/predecessor

func (bst *BST) Delete(value int) {
    bst.Root = deleteNode(bst.Root, value)
}

func deleteNode(node *Node, value int) *Node {
    if node == nil {
        return nil
    }
    
    if value < node.Value {
        node.Left = deleteNode(node.Left, value)
    } else if value > node.Value {
        node.Right = deleteNode(node.Right, value)
    } else {
        // Node to delete found
        if node.Left == nil {
            return node.Right
        }
        if node.Right == nil {
            return node.Left
        }
        
        // Two children: find inorder successor
        successor := findMin(node.Right)
        node.Value = successor.Value
        node.Right = deleteNode(node.Right, successor.Value)
    }
    return node
}
```

### Range Queries
```go
func (bst *BST) RangeSearch(min, max int) []int {
    var result []int
    rangeSearchHelper(bst.Root, min, max, &result)
    return result
}

func rangeSearchHelper(node *Node, min, max int, result *[]int) {
    if node == nil {
        return
    }
    
    if node.Value > min {
        rangeSearchHelper(node.Left, min, max, result)
    }
    
    if node.Value >= min && node.Value <= max {
        *result = append(*result, node.Value)
    }
    
    if node.Value < max {
        rangeSearchHelper(node.Right, min, max, result)
    }
}
```

## Practical Considerations

### When to Use BSTs
- **Dynamic Sets**: Frequent insertions/deletions with search requirements
- **Range Queries**: Finding elements within value ranges
- **Ordered Iteration**: Need sorted traversal of elements
- **Memory Constraints**: Lower overhead than hash tables

### When NOT to Use BSTs
- **Skewed Data**: Sorted or nearly sorted input sequences
- **High Performance**: Hash tables for simple key-value lookups
- **Guaranteed Balance**: Use self-balancing trees (AVL, Red-Black)
- **Concurrent Access**: Consider concurrent data structures

### Real-world Applications
- **Database Indexing**: B-trees (generalized BSTs)
- **Expression Parsing**: Syntax trees in compilers
- **File Systems**: Directory hierarchies
- **Priority Queues**: Binary heaps (complete binary trees)