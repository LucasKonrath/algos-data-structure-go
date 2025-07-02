# Three Sum Problem - Expert Analysis

## Problem Definition
Given an array of integers, find all unique triplets `[a, b, c]` such that `a + b + c = 0`. The solution must avoid duplicate triplets and handle edge cases efficiently.

## Algorithmic Approach
The implementation uses a **sorted array + two-pointer technique**, transforming the 3SUM problem into multiple 2SUM subproblems.

## Algorithm Breakdown

### Step 1: Array Sorting
```go
sort.Ints(nums)  // O(n log n)
```
**Purpose**: Enables two-pointer technique and simplifies duplicate handling

### Step 2: Fixed Element Iteration
```go
for i := 0; i < len(nums)-2; i++ {
    if i > 0 && nums[i] == nums[i-1] {
        continue // Skip duplicates for first element
    }
    // ... two-pointer search
}
```
**Strategy**: Fix the first element, solve 2SUM for remaining elements

### Step 3: Two-Pointer Search
```go
left, right := i+1, len(nums)-1

for left < right {
    sum := nums[i] + nums[left] + nums[right]
    if sum == 0 {
        // Found triplet, handle duplicates
    } else if sum < 0 {
        left++  // Need larger sum
    } else {
        right-- // Need smaller sum
    }
}
```

## Detailed Execution Trace

### Example Input: [-1, 0, 1, 2, -1, -4]

#### Phase 1: Sorting
```
Original: [-1, 0, 1, 2, -1, -4]
Sorted:   [-4, -1, -1, 0, 1, 2]
Indices:   0   1   2  3  4  5
```

#### Phase 2: Triplet Finding
```
i=0: nums[0] = -4, target = 4
  left=1(-1), right=5(2): sum=-4+(-1)+2=-3 < 0, left++
  left=2(-1), right=5(2): sum=-4+(-1)+2=-3 < 0, left++  
  left=3(0), right=5(2): sum=-4+0+2=-2 < 0, left++
  left=4(1), right=5(2): sum=-4+1+2=-1 < 0, left++
  left=5: left >= right, exit

i=1: nums[1] = -1, target = 1
  left=2(-1), right=5(2): sum=-1+(-1)+2=0 ✓
  Found triplet: [-1, -1, 2]
  Skip duplicates: left=3, right=4
  left=3(0), right=4(1): sum=-1+0+1=0 ✓  
  Found triplet: [-1, 0, 1]
  left=4, right=3: left >= right, exit

i=2: nums[2] = -1, duplicate of nums[1], skip

i=3: nums[3] = 0, target = 0
  left=4(1), right=5(2): sum=0+1+2=3 > 0, right--
  left=4, right=4: left >= right, exit

Result: [[-1, -1, 2], [-1, 0, 1]]
```

## Complexity Analysis

### Time Complexity: O(n²)
```
Sorting: O(n log n)
Outer loop: O(n) iterations
Inner two-pointer: O(n) per iteration
Total: O(n log n) + O(n²) = O(n²)
```

### Space Complexity: O(1) to O(n)
```
In-place sorting: O(1) extra space
Result storage: O(k) where k = number of triplets
Worst case k: O(n²) for dense solutions
```

## Duplicate Handling Strategy

### Three Levels of Deduplication

#### Level 1: First Element (i)
```go
if i > 0 && nums[i] == nums[i-1] {
    continue
}
```
**Prevents**: Duplicate triplets starting with same first element

#### Level 2: Second Element (left)
```go
for left < right && nums[left] == nums[left+1] {
    left++
}
```
**Prevents**: Duplicate triplets with same first and second elements

#### Level 3: Third Element (right)
```go
for left < right && nums[right] == nums[right-1] {
    right--
}
```
**Prevents**: Completely duplicate triplets

### Deduplication Example
```
Array: [-2, 0, 0, 2, 2]
Without deduplication: [[-2, 0, 2], [-2, 0, 2], [-2, 0, 2], [-2, 0, 2]]
With deduplication: [[-2, 0, 2]]
```

## Mathematical Analysis

### Problem Transformation
```
3SUM: Find (a, b, c) where a + b + c = 0
Transform to: Find (b, c) where b + c = -a

For each fixed a, solve 2SUM on remaining elements
```

### Combinatorial Bounds
```
Maximum triplets for array of n elements:
- Theoretical maximum: C(n,3) = n(n-1)(n-2)/6
- Practical constraint: Target sum limits actual count
- Average case: Much less than theoretical maximum
```

### Two-Pointer Correctness
**Invariant**: At each step, all triplets with sum ≠ 0 involving elements outside [left, right] have been considered.

**Proof**: 
- If sum < 0: nums[left] too small, increment left
- If sum > 0: nums[right] too large, decrement right  
- If sum = 0: Found valid triplet, move both pointers

## Optimization Techniques

### Early Termination
```go
if nums[i] > 0 {
    break // All remaining elements positive
}

if nums[i] + nums[len(nums)-2] + nums[len(nums)-1] < 0 {
    continue // Impossible to reach 0 with this i
}

if nums[i] + nums[i+1] + nums[i+2] > 0 {
    break // All remaining combinations > 0
}
```

### Hash Set Alternative (Higher Space)
```go
func threeSumHashSet(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        seen := make(map[int]bool)
        target := -nums[i]
        
        for j := i + 1; j < len(nums); j++ {
            complement := target - nums[j]
            if seen[complement] {
                result = append(result, []int{nums[i], complement, nums[j]})
                // Skip duplicates
                for j+1 < len(nums) && nums[j] == nums[j+1] {
                    j++
                }
            }
            seen[nums[j]] = true
        }
    }
    return result
}
```

## Variant Problems

### 3Sum Closest
```go
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]
    
    for i := 0; i < len(nums)-2; i++ {
        left, right := i+1, len(nums)-1
        
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if abs(target-sum) < abs(target-closest) {
                closest = sum
            }
            
            if sum < target {
                left++
            } else {
                right--
            }
        }
    }
    return closest
}
```

### 4Sum Extension
```go
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    
    for i := 0; i < len(nums)-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        for j := i + 1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            
            // Apply 2Sum on remaining elements
            left, right := j+1, len(nums)-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                // ... similar logic
            }
        }
    }
    return result
}
```

## Performance Characteristics

### Best Case: O(n log n)
- Few or no valid triplets
- Heavy pruning from early termination

### Average Case: O(n²)
- Balanced distribution of elements
- Moderate number of valid triplets

### Worst Case: O(n²)
- Dense solution space
- Many duplicate handling operations

### Cache Performance
- **Sequential Access**: Good cache locality during sorting
- **Random Access**: Two-pointer technique has good spatial locality
- **Memory Bandwidth**: Minimal due to in-place operations

## Practical Considerations

### Input Constraints
```go
func validateInput(nums []int) error {
    if len(nums) < 3 {
        return errors.New("array must have at least 3 elements")
    }
    return nil
}
```

### Overflow Handling
```go
func safeSum(a, b, c int) (int, bool) {
    if a > 0 && b > 0 && c > math.MaxInt-a-b {
        return 0, false // Overflow
    }
    if a < 0 && b < 0 && c < math.MinInt-a-b {
        return 0, false // Underflow
    }
    return a + b + c, true
}
```

### Generalization to K-Sum
```go
func kSum(nums []int, k int, target int) [][]int {
    if k == 2 {
        return twoSum(nums, target)
    }
    
    result := [][]int{}
    for i := 0; i < len(nums)-k+1; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        subResults := kSum(nums[i+1:], k-1, target-nums[i])
        for _, sub := range subResults {
            result = append(result, append([]int{nums[i]}, sub...))
        }
    }
    return result
}
```

## Real-world Applications
- **Computational Geometry**: Finding collinear points
- **Chemistry**: Balancing chemical equations
- **Finance**: Portfolio optimization with constraints
- **Game Development**: Physics constraint solving
- **Machine Learning**: Feature selection with sum constraints