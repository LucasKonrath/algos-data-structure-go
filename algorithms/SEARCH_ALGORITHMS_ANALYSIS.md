# Search Algorithms - Expert Analysis

## Overview
Search algorithms are fundamental techniques for locating specific elements within data structures. This analysis covers two primary approaches: **Linear Search** and **Binary Search**, examining their trade-offs, optimizations, and practical applications.

# Linear Search Algorithm

## Algorithm Overview
Linear Search is the simplest search algorithm that examines each element sequentially until the target is found or the end of the array is reached. It makes no assumptions about data ordering.

## Implementation Analysis

### Core Algorithm
```go
func linearSearch(arr []int, target int) int {
    for i, v := range arr {
        if v == target {
            return i  // Return index of first occurrence
        }
    }
    return -1  // Element not found
}
```

### Algorithm Properties
- **Sequential Access**: Examines elements one by one
- **No Preprocessing**: Works on unsorted data
- **First Match**: Returns index of first occurrence
- **Termination**: Stops immediately when found

## Complexity Analysis

### Time Complexity
```
Best Case: O(1)
- Target is the first element
- Single comparison required

Average Case: O(n/2) = O(n)
- Target equally likely to be anywhere
- Expected comparisons: (n+1)/2

Worst Case: O(n)
- Target is last element or not present
- All n elements must be examined
```

### Space Complexity
```
Space: O(1)
- Only uses constant extra space
- In-place algorithm with index variable
```

## Linear Search Variants

### Multiple Occurrences
```go
func linearSearchAll(arr []int, target int) []int {
    indices := []int{}
    for i, v := range arr {
        if v == target {
            indices = append(indices, i)
        }
    }
    return indices
}
```

### Early Termination with Sentinel
```go
func linearSearchSentinel(arr []int, target int) int {
    n := len(arr)
    if n == 0 {
        return -1
    }
    
    // Store last element and place sentinel
    last := arr[n-1]
    arr[n-1] = target
    
    i := 0
    for arr[i] != target {
        i++
    }
    
    // Restore last element
    arr[n-1] = last
    
    // Check if found or sentinel reached
    if i < n-1 || arr[n-1] == target {
        return i
    }
    return -1
}
```

### Transposition Heuristic
```go
func linearSearchTranspose(arr []int, target int) int {
    for i, v := range arr {
        if v == target {
            if i > 0 {
                // Move found element one position forward
                arr[i], arr[i-1] = arr[i-1], arr[i]
                return i - 1
            }
            return i
        }
    }
    return -1
}
```

# Binary Search Algorithm

## Algorithm Overview
Binary Search is an efficient search algorithm for sorted arrays that repeatedly divides the search space in half, eliminating half of the remaining elements with each comparison.

## Implementation Analysis

### Core Algorithm
```go
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2  // Overflow-safe midpoint
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1   // Search right half
        } else {
            right = mid - 1  // Search left half
        }
    }
    return -1  // Element not found
}
```

### Key Implementation Details
1. **Overflow Prevention**: `mid = left + (right-left)/2` instead of `(left+right)/2`
2. **Inclusive Bounds**: Both left and right are valid indices
3. **Loop Condition**: `left <= right` ensures all elements are checked
4. **Boundary Updates**: `mid+1` and `mid-1` to avoid infinite loops

## Detailed Execution Trace

### Example: Search for 7 in [1, 3, 5, 7, 9, 11, 13]

```
Array: [1, 3, 5, 7, 9, 11, 13]
Index:  0  1  2  3  4   5   6

Iteration 1:
left=0, right=6, mid=3
arr[3]=7, target=7 → Found! Return 3

Single iteration due to lucky middle element selection.

Alternative example - Search for 11:

Iteration 1:
left=0, right=6, mid=3
arr[3]=7 < 11 → Search right half
left=4, right=6

Iteration 2:
left=4, right=6, mid=5
arr[5]=11 == 11 → Found! Return 5
```

### Unsuccessful Search Example
```
Search for 8 in [1, 3, 5, 7, 9, 11, 13]:

Iteration 1: left=0, right=6, mid=3, arr[3]=7 < 8 → left=4
Iteration 2: left=4, right=6, mid=5, arr[5]=11 > 8 → right=4
Iteration 3: left=4, right=4, mid=4, arr[4]=9 > 8 → right=3
Iteration 4: left=4, right=3 → left > right, return -1
```

## Complexity Analysis

### Time Complexity: O(log n)
```
Mathematical Analysis:
After k iterations, search space ≤ n/2^k
Algorithm terminates when n/2^k < 1
Therefore: k > log₂(n)
Maximum iterations: ⌈log₂(n)⌉

Recurrence Relation:
T(n) = T(n/2) + O(1)
T(n) = O(log n) by Master Theorem
```

### Space Complexity
```
Iterative: O(1) - Only uses constant extra variables
Recursive: O(log n) - Call stack depth equals tree height
```

## Binary Search Variants

### Recursive Implementation
```go
func binarySearchRecursive(arr []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := left + (right-left)/2
    
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        return binarySearchRecursive(arr, target, mid+1, right)
    } else {
        return binarySearchRecursive(arr, target, left, mid-1)
    }
}
```

### Lower Bound (First Occurrence)
```go
func lowerBound(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
```

### Upper Bound (Last Occurrence + 1)
```go
func upperBound(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        if arr[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
```

### Range Search
```go
func searchRange(arr []int, target int) [2]int {
    first := lowerBound(arr, target)
    if first == len(arr) || arr[first] != target {
        return [2]int{-1, -1}
    }
    
    last := upperBound(arr, target) - 1
    return [2]int{first, last}
}
```

## Advanced Binary Search Applications

### Search in Rotated Array
```go
func searchRotated(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        }
        
        // Determine which half is sorted
        if arr[left] <= arr[mid] {
            // Left half is sorted
            if target >= arr[left] && target < arr[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right half is sorted
            if target > arr[mid] && target <= arr[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

### Search 2D Matrix
```go
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    
    rows, cols := len(matrix), len(matrix[0])
    left, right := 0, rows*cols-1
    
    for left <= right {
        mid := left + (right-left)/2
        midValue := matrix[mid/cols][mid%cols]
        
        if midValue == target {
            return true
        } else if midValue < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

### Find Peak Element
```go
func findPeakElement(arr []int) int {
    left, right := 0, len(arr)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] > arr[mid+1] {
            // Peak is in left half (including mid)
            right = mid
        } else {
            // Peak is in right half
            left = mid + 1
        }
    }
    return left
}
```

## Performance Comparison

### Theoretical Analysis
```
                Linear Search    Binary Search
Time (Best)         O(1)            O(1)
Time (Average)      O(n)            O(log n)
Time (Worst)        O(n)            O(log n)
Space               O(1)            O(1) iterative
Preprocessing       None            O(n log n) sorting
Data Requirement    None            Must be sorted
```

### Breakeven Analysis
```go
// Calculate when binary search becomes worthwhile
func breakevenPoint() {
    // Binary search: log₂(n) + sorting cost
    // Linear search: n/2 average comparisons
    
    // For repeated searches on same data:
    // k * log₂(n) vs k * n/2
    // Binary search wins when k > 1 for large n
    
    // For single search:
    // n*log₂(n) + log₂(n) vs n/2
    // Linear search often better for single queries
}
```

## Cache Performance Analysis

### Linear Search Cache Behavior
```
Cache Performance: Excellent
- Sequential memory access
- High spatial locality  
- Prefetcher-friendly access pattern
- Cache lines fully utilized

Memory Access Pattern:
addr[0], addr[1], addr[2], ... (sequential)
```

### Binary Search Cache Behavior
```
Cache Performance: Poor for large arrays
- Random memory access pattern
- Low spatial locality
- Poor cache line utilization
- Each iteration may cause cache miss

Memory Access Pattern:
addr[n/2], addr[n/4] or addr[3n/4], ... (scattered)
```

### Cache-Friendly Binary Search
```go
func cacheOptimizedBinarySearch(arr []int, target int) int {
    const CACHE_LINE_SIZE = 64  // bytes
    const INTS_PER_CACHE_LINE = CACHE_LINE_SIZE / 8
    
    // Switch to linear search for small subarrays
    if len(arr) <= INTS_PER_CACHE_LINE {
        return linearSearch(arr, target)
    }
    
    // Standard binary search for larger arrays
    return binarySearch(arr, target)
}
```

## Interpolation Search
```go
func interpolationSearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right && target >= arr[left] && target <= arr[right] {
        if left == right {
            if arr[left] == target {
                return left
            }
            return -1
        }
        
        // Interpolation formula
        pos := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])
        
        if arr[pos] == target {
            return pos
        } else if arr[pos] < target {
            left = pos + 1
        } else {
            right = pos - 1
        }
    }
    return -1
}
```

## Exponential Search
```go
func exponentialSearch(arr []int, target int) int {
    n := len(arr)
    if n == 0 {
        return -1
    }
    
    if arr[0] == target {
        return 0
    }
    
    // Find range for binary search
    bound := 1
    for bound < n && arr[bound] < target {
        bound *= 2
    }
    
    // Binary search in found range
    left := bound / 2
    right := min(bound, n-1)
    
    return binarySearchRange(arr, target, left, right)
}
```

## Error Handling and Edge Cases

### Robust Implementations
```go
func binarySearchSafe(arr []int, target int) (int, error) {
    if arr == nil {
        return -1, errors.New("array cannot be nil")
    }
    
    if len(arr) == 0 {
        return -1, nil  // Empty array, not found
    }
    
    // Verify array is sorted
    if !isSorted(arr) {
        return -1, errors.New("array must be sorted for binary search")
    }
    
    return binarySearch(arr, target), nil
}

func isSorted(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i-1] {
            return false
        }
    }
    return true
}
```

### Integer Overflow Protection
```go
func safeMidpoint(left, right int) int {
    // Prevent overflow in (left + right) / 2
    return left + (right-left)/2
}

// Alternative using bit shifting
func midpointBitShift(left, right int) int {
    return (left + right) >> 1  // Only safe if both are positive
}
```

## Generic Implementations

### Type-Agnostic Search
```go
func BinarySearchGeneric[T any](arr []T, target T, compare func(T, T) int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        cmp := compare(arr[mid], target)
        
        if cmp == 0 {
            return mid
        } else if cmp < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

// Usage example
func searchStrings() {
    words := []string{"apple", "banana", "cherry", "date"}
    index := BinarySearchGeneric(words, "cherry", strings.Compare)
}
```

## Real-World Applications

### Database Indexing
```go
// B-tree search in database index
func btreeSearch(node *BTreeNode, key int) *Record {
    // Binary search within node
    pos := binarySearch(node.keys, key)
    
    if pos >= 0 {
        return node.records[pos]
    }
    
    if node.isLeaf {
        return nil
    }
    
    // Recursively search appropriate child
    childIndex := findChildIndex(node, key)
    return btreeSearch(node.children[childIndex], key)
}
```

### System Call Optimization
```go
// Search system call tables
func findSystemCall(syscallNumber int) *SyscallHandler {
    // Binary search in sorted syscall table
    index := binarySearch(syscallTable.numbers, syscallNumber)
    if index >= 0 {
        return syscallTable.handlers[index]
    }
    return nil
}
```

### Network Routing
```go
// IP address lookup in routing table
func longestPrefixMatch(routingTable []Route, destIP uint32) *Route {
    // Binary search for longest matching prefix
    left, right := 0, len(routingTable)-1
    bestMatch := (*Route)(nil)
    
    for left <= right {
        mid := left + (right-left)/2
        route := &routingTable[mid]
        
        if route.matches(destIP) {
            bestMatch = route
            left = mid + 1  // Look for longer prefix
        } else if route.network < destIP {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return bestMatch
}
```

## Testing and Verification

### Comprehensive Test Suite
```go
func TestSearchAlgorithms(t *testing.T) {
    testCases := []struct {
        name   string
        arr    []int
        target int
        want   int
    }{
        {"Empty array", []int{}, 5, -1},
        {"Single element found", []int{5}, 5, 0},
        {"Single element not found", []int{3}, 5, -1},
        {"First element", []int{1, 2, 3, 4, 5}, 1, 0},
        {"Last element", []int{1, 2, 3, 4, 5}, 5, 4},
        {"Middle element", []int{1, 2, 3, 4, 5}, 3, 2},
        {"Not found", []int{1, 2, 3, 4, 5}, 6, -1},
        {"Duplicates", []int{1, 2, 2, 2, 5}, 2, 1}, // Returns first occurrence
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Test both algorithms
            linearResult := linearSearch(tc.arr, tc.target)
            binaryResult := binarySearch(tc.arr, tc.target)
            
            if linearResult != tc.want {
                t.Errorf("Linear search: got %d, want %d", linearResult, tc.want)
            }
            
            if binaryResult != tc.want {
                t.Errorf("Binary search: got %d, want %d", binaryResult, tc.want)
            }
        })
    }
}
```

## Performance Benchmarking
```go
func BenchmarkSearchAlgorithms(b *testing.B) {
    sizes := []int{100, 1000, 10000, 100000}
    
    for _, size := range sizes {
        arr := make([]int, size)
        for i := range arr {
            arr[i] = i * 2  // Even numbers, sorted
        }
        
        target := size  // Element not in array (worst case)
        
        b.Run(fmt.Sprintf("Linear/%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                linearSearch(arr, target)
            }
        })
        
        b.Run(fmt.Sprintf("Binary/%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                binarySearch(arr, target)
            }
        })
    }
}
```

The choice between linear and binary search depends on factors including data size, search frequency, memory constraints, and whether the data is already sorted. Understanding these trade-offs is crucial for optimal algorithm selection in real-world applications.