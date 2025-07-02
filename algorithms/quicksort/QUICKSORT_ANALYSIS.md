# QuickSort Algorithm - Expert Analysis

## Algorithm Overview
QuickSort is a **divide-and-conquer** sorting algorithm that works by selecting a 'pivot' element and partitioning the array around it. Elements smaller than the pivot go to the left, larger elements go to the right.

## Implementation Analysis

### Core Components
1. **Pivot Selection**: Uses last element as pivot (Lomuto partition scheme)
2. **Partitioning**: Two-pointer technique for in-place partitioning
3. **Recursive Calls**: Divide-and-conquer on sub-arrays

### Partitioning Mechanism
```
Initial Array: [3, 6, 8, 10, 1, 2, 1]
Pivot: 1 (last element)

Step-by-step partitioning:
i = -1, j = 0 to 5

j=0: arr[0]=3 > pivot=1, no swap, i=-1
j=1: arr[1]=6 > pivot=1, no swap, i=-1  
j=2: arr[2]=8 > pivot=1, no swap, i=-1
j=3: arr[3]=10 > pivot=1, no swap, i=-1
j=4: arr[4]=1 <= pivot=1, i=0, swap arr[0] and arr[4]
j=5: arr[5]=2 > pivot=1, no swap, i=0

Final swap: arr[1] and arr[6] (pivot position)
Result: [1, 6, 8, 10, 3, 2, 1] → [1, 1, 8, 10, 3, 2, 6]
```

## Complexity Analysis

### Time Complexity
- **Best Case**: O(n log n) - Balanced partitions
- **Average Case**: O(n log n) - Random pivot selection
- **Worst Case**: O(n²) - Already sorted array with poor pivot

### Space Complexity
- **Best/Average**: O(log n) - Recursion stack depth
- **Worst Case**: O(n) - Unbalanced recursion tree

## Partition Analysis

### Lomuto Partition Scheme
```go
func partition(arr []int, low int, high int) int {
    pivot := arr[high]  // Last element as pivot
    i := low - 1        // Index of smaller element
    
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
```

### Invariants Maintained
1. **Elements [low...i]**: ≤ pivot
2. **Elements [i+1...j-1]**: > pivot  
3. **Elements [j...high-1]**: Unprocessed
4. **Element [high]**: Pivot

## Performance Characteristics

### Pivot Selection Impact
- **Last Element**: Simple but vulnerable to sorted inputs
- **Random Pivot**: Better average case performance
- **Median-of-Three**: Improved worst-case behavior

### Cache Performance
- **Locality**: Good due to sequential access during partitioning
- **Memory Access**: In-place sorting minimizes memory overhead

## Optimization Opportunities

### Tail Recursion Elimination
```go
func quicksortOptimized(arr []int, low int, high int) {
    for low < high {
        pi := partition(arr, low, high)
        
        // Recurse on smaller partition, iterate on larger
        if pi - low < high - pi {
            quicksort(arr, low, pi-1)
            low = pi + 1
        } else {
            quicksort(arr, pi+1, high)
            high = pi - 1
        }
    }
}
```

### Hybrid Approaches
- **Introsort**: Switch to HeapSort for worst-case scenarios
- **Insertion Sort**: Use for small arrays (< 10 elements)
- **Dual-Pivot**: Java's approach for better performance

## Mathematical Properties

### Recurrence Relations
```
T(n) = T(k) + T(n-k-1) + Θ(n)
```
Where k is the number of elements smaller than pivot.

### Expected Performance
```
E[T(n)] = (2/n) * Σ(k=0 to n-1) E[T(k)] + Θ(n)
```
Solving gives E[T(n)] = O(n log n)

## Stability and Adaptability
- **Stability**: Not stable (relative order of equal elements may change)
- **Adaptability**: Not adaptive (doesn't benefit from pre-sorted input)
- **Online**: Not online (requires entire array)

## Implementation Notes
- Uses in-place partitioning for O(1) extra space
- Recursive implementation creates O(log n) stack frames
- Vulnerable to adversarial inputs without randomization