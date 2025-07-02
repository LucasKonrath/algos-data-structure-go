# Recursion Algorithms - Comprehensive Analysis

## Table of Contents
1. [Recursion Fundamentals](#recursion-fundamentals)
2. [Factorial Algorithm](#factorial-algorithm)
3. [Fibonacci Variants](#fibonacci-variants)
4. [Dynamic Programming Knapsack](#dynamic-programming-knapsack)
5. [N-Queens Problem](#n-queens-problem)
6. [Permutation Generation](#permutation-generation)
7. [Mathematical Foundations](#mathematical-foundations)
8. [Performance Analysis](#performance-analysis)
9. [Real-World Applications](#real-world-applications)
10. [Optimization Techniques](#optimization-techniques)

## Recursion Fundamentals

### Core Principles
Recursion is a programming technique where a function calls itself to solve smaller instances of the same problem. Every recursive algorithm must have:

1. **Base Case**: A condition that stops the recursion
2. **Recursive Case**: The function calls itself with modified parameters
3. **Progress**: Each recursive call must move closer to the base case

### Types of Recursion
- **Linear Recursion**: Each function call makes at most one recursive call
- **Tree Recursion**: Each function call makes multiple recursive calls
- **Tail Recursion**: The recursive call is the last operation in the function

### Recursion vs Iteration Trade-offs
| Aspect | Recursion | Iteration |
|--------|-----------|-----------|
| Memory | O(n) stack space | O(1) stack space |
| Readability | Often cleaner | Can be more verbose |
| Performance | Function call overhead | Direct loop execution |
| Stack Overflow | Possible with deep recursion | Not applicable |

## Factorial Algorithm

### Implementation Analysis

```go
func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

### Algorithm Breakdown

#### Base Cases
- `n == 0`: Returns 1 (0! = 1 by definition)
- `n == 1`: Returns 1 (1! = 1)

#### Recursive Case
- `n > 1`: Returns `n * factorial(n-1)`

### Complexity Analysis
- **Time Complexity**: O(n) - makes n recursive calls
- **Space Complexity**: O(n) - n stack frames
- **Recurrence Relation**: T(n) = T(n-1) + O(1)

### Execution Trace Example
**Input**: n = 5

```
factorial(5)
├── 5 * factorial(4)
    ├── 4 * factorial(3)
        ├── 3 * factorial(2)
            ├── 2 * factorial(1)
                └── 1 (base case)
            └── 2 * 1 = 2
        └── 3 * 2 = 6
    └── 4 * 6 = 24
└── 5 * 24 = 120

Call Stack Depth: 5
Total Multiplications: 4
Result: 120
```

### Stack Frame Analysis
```
Stack Frame 1: factorial(5) - n=5, waiting for factorial(4)
Stack Frame 2: factorial(4) - n=4, waiting for factorial(3)
Stack Frame 3: factorial(3) - n=3, waiting for factorial(2)
Stack Frame 4: factorial(2) - n=2, waiting for factorial(1)
Stack Frame 5: factorial(1) - n=1, returns 1
```

### Optimization: Tail Recursion
```go
func factorialTailRec(n int, acc int) int {
    if n == 0 || n == 1 {
        return acc
    }
    return factorialTailRec(n-1, n*acc)
}

func factorial(n int) int {
    return factorialTailRec(n, 1)
}
```

## Fibonacci Variants

### 1. Naive Recursive Implementation

```go
func fibonacci(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 || n == 2 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

#### Complexity Analysis
- **Time Complexity**: O(2^n) - exponential due to repeated calculations
- **Space Complexity**: O(n) - maximum recursion depth

#### Execution Trace for fibonacci(5)
```
fibonacci(5)
├── fibonacci(4)
│   ├── fibonacci(3)
│   │   ├── fibonacci(2) → 1
│   │   └── fibonacci(1) → 1
│   │   └── result: 2
│   └── fibonacci(2) → 1
│   └── result: 3
└── fibonacci(3)
    ├── fibonacci(2) → 1
    └── fibonacci(1) → 1
    └── result: 2
└── result: 5

Total Function Calls: 15
Redundant Calculations: fibonacci(3) computed twice, fibonacci(2) computed 3 times
```

### 2. Memoized Recursive Implementation

```go
func fibonacci(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
    return memo[n]
}
```

#### Complexity Analysis
- **Time Complexity**: O(n) - each subproblem solved once
- **Space Complexity**: O(n) - memoization table + recursion stack

#### Memoization Benefits
- Eliminates redundant calculations
- Transforms exponential time to linear time
- Maintains recursive structure for readability

### 3. Tabulation (Bottom-Up DP)

```go
func fibonacciTabulation(n int) int {
    if n <= 1 {
        return n
    }
    
    fibTable := make([]int, n+1)
    fibTable[0], fibTable[1] = 0, 1
    
    for i := 2; i <= n; i++ {
        fibTable[i] = fibTable[i-1] + fibTable[i-2]
    }
    
    return fibTable[n]
}
```

#### Complexity Analysis
- **Time Complexity**: O(n) - single loop
- **Space Complexity**: O(n) - table storage

#### Space-Optimized Version
```go
func fibonacciOptimized(n int) int {
    if n <= 1 {
        return n
    }
    
    prev2, prev1 := 0, 1
    for i := 2; i <= n; i++ {
        current := prev1 + prev2
        prev2, prev1 = prev1, current
    }
    return prev1
}
```
- **Space Complexity**: O(1) - only tracking last two values

## Dynamic Programming Knapsack

### 1. Memoized Recursive Implementation

```go
func knapsack(weights []int, values []int, capacity int, n int, memo map[string]int) int {
    if n == 0 || capacity == 0 {
        return 0
    }
    
    key := fmt.Sprintf("%d,%d", n, capacity)
    if val, exists := memo[key]; exists {
        return val
    }
    
    if weights[n-1] > capacity {
        memo[key] = knapsack(weights, values, capacity, n-1, memo)
        return memo[key]
    } else {
        includeItem := values[n-1] + knapsack(weights, values, capacity-weights[n-1], n-1, memo)
        excludeItem := knapsack(weights, values, capacity, n-1, memo)
        memo[key] = max(includeItem, excludeItem)
        return memo[key]
    }
}
```

#### Algorithm Logic
1. **Base Case**: No items or no capacity → return 0
2. **Memoization Check**: Return cached result if available
3. **Decision Making**:
   - If item doesn't fit: exclude it
   - If item fits: choose maximum of including or excluding

#### Complexity Analysis
- **Time Complexity**: O(n × capacity) - each subproblem solved once
- **Space Complexity**: O(n × capacity) - memoization table + O(n) recursion stack

### 2. Tabulation Implementation

```go
func knapsackTabulation(weights []int, values []int, capacity int) int {
    n := len(weights)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, capacity+1)
    }
    
    for i := 1; i <= n; i++ {
        for w := 0; w <= capacity; w++ {
            if weights[i-1] <= w {
                dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
            } else {
                dp[i][w] = dp[i-1][w]
            }
        }
    }
    
    return dp[n][capacity]
}
```

#### DP Table Construction
The table `dp[i][w]` represents the maximum value achievable with first `i` items and capacity `w`.

#### Example Execution
**Input**: weights = [1, 2, 3], values = [6, 10, 12], capacity = 5

```
DP Table Construction:
    w=0  w=1  w=2  w=3  w=4  w=5
i=0  0    0    0    0    0    0
i=1  0    6    6    6    6    6    (item 1: weight=1, value=6)
i=2  0    6   10   16   16   16    (item 2: weight=2, value=10)
i=3  0    6   10   16   18   22    (item 3: weight=3, value=12)

Final Answer: dp[3][5] = 22
Optimal Selection: Items 2 and 3 (values 10 + 12 = 22, weights 2 + 3 = 5)
```

## N-Queens Problem

### Implementation Analysis

```go
func solveNQueens(n int) [][]string {
    solutions := [][]string{}
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }
    
    var backtrack func(row int)
    backtrack = func(row int) {
        if row == n {
            solution := make([]string, n)
            for i, row := range board {
                solution[i] = string(row)
            }
            solutions = append(solutions, solution)
        }
        
        for col := 0; col < n; col++ {
            if isSafe(board, row, col, n) {
                board[row][col] = 'Q'
                backtrack(row + 1)
                board[row][col] = '.'
            }
        }
    }
    backtrack(0)
    return solutions
}
```

### Algorithm Breakdown

#### Backtracking Strategy
1. **Place Queen**: Try placing a queen in each column of current row
2. **Check Safety**: Verify no conflicts with existing queens
3. **Recurse**: Move to next row if placement is safe
4. **Backtrack**: Remove queen and try next position if dead end

#### Safety Check Implementation
```go
func isSafe(board [][]byte, row, col, n int) bool {
    // Check column
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    
    // Check diagonal (top-left to bottom-right)
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    
    // Check diagonal (top-right to bottom-left)
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    
    return true
}
```

### Complexity Analysis
- **Time Complexity**: O(n!) - in worst case, try all permutations
- **Space Complexity**: O(n²) - board storage + O(n) recursion depth
- **Pruning Effect**: Constraint propagation reduces actual branching factor

### Execution Trace for N=4
```
Initial Board:
. . . .
. . . .
. . . .
. . . .

Step 1: Try row 0
├── Col 0: Place Q
│   Q . . .
│   . . . .
│   . . . .
│   . . . .
│   
│   Step 2: Try row 1
│   ├── Col 0: Unsafe (same column)
│   ├── Col 1: Unsafe (diagonal)
│   ├── Col 2: Place Q
│   │   Q . . .
│   │   . . Q .
│   │   . . . .
│   │   . . . .
│   │   
│   │   Step 3: Try row 2
│   │   ├── Col 0: Unsafe (same column)
│   │   ├── Col 1: Place Q
│   │   │   Q . . .
│   │   │   . . Q .
│   │   │   . Q . .
│   │   │   . . . .
│   │   │   
│   │   │   Step 4: Try row 3
│   │   │   ├── Col 0: Unsafe (same column)
│   │   │   ├── Col 1: Unsafe (same column)
│   │   │   ├── Col 2: Unsafe (same column)
│   │   │   └── Col 3: Place Q → Solution Found!
│   │   │       Q . . .
│   │   │       . . Q .
│   │   │       . Q . .
│   │   │       . . . Q

Solution 1: [".Q..", "...Q", "Q...", "..Q."]
```

### Optimization Techniques
1. **Bitmasking**: Use integers to represent column and diagonal occupancy
2. **Symmetry Breaking**: Only search half the solutions and mirror them
3. **Constraint Propagation**: More sophisticated pruning

## Permutation Generation

### Implementation Analysis

```go
func permute(nums []int) [][]int {
    var result [][]int
    var backtrack func([]int, []int)
    
    backtrack = func(current []int, remaining []int) {
        if len(remaining) == 0 {
            result = append(result, append([]int{}, current...))
            return
        }
        
        for i, num := range remaining {
            newCurrent := append(current, num)
            newRemaining := append(append([]int{}, remaining[:i]...), remaining[i+1:]...)
            backtrack(newCurrent, newRemaining)
        }
    }
    
    backtrack([]int{}, nums)
    return result
}
```

### Algorithm Breakdown

#### Backtracking Strategy
1. **Base Case**: No remaining elements → add current permutation to result
2. **Recursive Case**: For each remaining element:
   - Add it to current permutation
   - Create new remaining list without selected element
   - Recurse with updated parameters

#### Complexity Analysis
- **Time Complexity**: O(n! × n) - n! permutations, each taking O(n) to generate
- **Space Complexity**: O(n! × n) - store all permutations + O(n) recursion depth

### Execution Trace for [1, 2, 3]
```
permute([1, 2, 3])
├── current=[], remaining=[1, 2, 3]
│   ├── Pick 1: current=[1], remaining=[2, 3]
│   │   ├── Pick 2: current=[1, 2], remaining=[3]
│   │   │   └── Pick 3: current=[1, 2, 3], remaining=[] → Add [1, 2, 3]
│   │   └── Pick 3: current=[1, 3], remaining=[2]
│   │       └── Pick 2: current=[1, 3, 2], remaining=[] → Add [1, 3, 2]
│   ├── Pick 2: current=[2], remaining=[1, 3]
│   │   ├── Pick 1: current=[2, 1], remaining=[3]
│   │   │   └── Pick 3: current=[2, 1, 3], remaining=[] → Add [2, 1, 3]
│   │   └── Pick 3: current=[2, 3], remaining=[1]
│   │       └── Pick 1: current=[2, 3, 1], remaining=[] → Add [2, 3, 1]
│   └── Pick 3: current=[3], remaining=[1, 2]
│       ├── Pick 1: current=[3, 1], remaining=[2]
│       │   └── Pick 2: current=[3, 1, 2], remaining=[] → Add [3, 1, 2]
│       └── Pick 2: current=[3, 2], remaining=[1]
│           └── Pick 1: current=[3, 2, 1], remaining=[] → Add [3, 2, 1]

Result: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
```

### Alternative Implementation (In-Place)
```go
func permuteInPlace(nums []int) [][]int {
    result := [][]int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {
        if start == len(nums) {
            result = append(result, append([]int{}, nums...))
            return
        }
        
        for i := start; i < len(nums); i++ {
            nums[start], nums[i] = nums[i], nums[start]  // Swap
            backtrack(start + 1)
            nums[start], nums[i] = nums[i], nums[start]  // Backtrack
        }
    }
    
    backtrack(0)
    return result
}
```

## Mathematical Foundations

### Recurrence Relations

#### Factorial
- **Recurrence**: T(n) = T(n-1) + O(1)
- **Solution**: T(n) = O(n)
- **Generating Function**: F(x) = ∑(n! × xⁿ)

#### Fibonacci
- **Recurrence**: F(n) = F(n-1) + F(n-2)
- **Characteristic Equation**: x² = x + 1
- **Closed Form**: F(n) = (φⁿ - ψⁿ)/√5 where φ = (1+√5)/2, ψ = (1-√5)/2

#### Knapsack
- **Recurrence**: K(n,w) = max(K(n-1,w), K(n-1,w-wₙ) + vₙ)
- **Optimal Substructure**: Optimal solution contains optimal solutions to subproblems

### Proof Techniques

#### Strong Induction for Factorial
**Base Case**: factorial(0) = 1 = 0! ✓
**Inductive Step**: Assume factorial(k) = k! for all k ≤ n-1
**Prove**: factorial(n) = n!
**Proof**: factorial(n) = n × factorial(n-1) = n × (n-1)! = n! ✓

#### Correctness of N-Queens Backtracking
**Invariant**: At each level, all placed queens are in safe positions
**Termination**: Either all queens are placed (solution found) or all possibilities exhausted
**Completeness**: Algorithm explores all valid configurations

## Performance Analysis

### Time Complexity Summary
| Algorithm | Best Case | Average Case | Worst Case |
|-----------|-----------|--------------|------------|
| Factorial | O(n) | O(n) | O(n) |
| Fibonacci (Naive) | O(2ⁿ) | O(2ⁿ) | O(2ⁿ) |
| Fibonacci (Memo) | O(n) | O(n) | O(n) |
| Fibonacci (Tabulation) | O(n) | O(n) | O(n) |
| Knapsack (Memo) | O(n×W) | O(n×W) | O(n×W) |
| Knapsack (Tabulation) | O(n×W) | O(n×W) | O(n×W) |
| N-Queens | O(n!) | O(n!) | O(n!) |
| Permutations | O(n!) | O(n!) | O(n!) |

### Space Complexity Analysis
| Algorithm | Auxiliary Space | Stack Space | Total |
|-----------|----------------|-------------|-------|
| Factorial | O(1) | O(n) | O(n) |
| Fibonacci (Naive) | O(1) | O(n) | O(n) |
| Fibonacci (Memo) | O(n) | O(n) | O(n) |
| Fibonacci (Tabulation) | O(n) | O(1) | O(n) |
| Knapsack (Memo) | O(n×W) | O(n) | O(n×W) |
| Knapsack (Tabulation) | O(n×W) | O(1) | O(n×W) |
| N-Queens | O(n²) | O(n) | O(n²) |
| Permutations | O(n!) | O(n) | O(n!) |

### Empirical Performance Testing
```go
func benchmarkFibonacci(n int) {
    // Naive
    start := time.Now()
    result1 := fibonacciNaive(n)
    naiveTime := time.Since(start)
    
    // Memoized
    start = time.Now()
    memo := make(map[int]int)
    result2 := fibonacciMemo(n, memo)
    memoTime := time.Since(start)
    
    // Tabulation
    start = time.Now()
    result3 := fibonacciTabulation(n)
    tabulationTime := time.Since(start)
    
    fmt.Printf("n=%d: Naive=%v, Memo=%v, Tabulation=%v\n", 
               n, naiveTime, memoTime, tabulationTime)
}
```

## Real-World Applications

### Factorial Applications
1. **Combinatorics**: Permutation and combination calculations
2. **Probability**: Calculating probabilities in discrete distributions
3. **Cryptography**: Key generation and encryption algorithms
4. **Computer Graphics**: Bezier curve calculations

### Fibonacci Applications
1. **Financial Markets**: Fibonacci retracements in trading
2. **Nature**: Spiral patterns in shells, flowers, and galaxies
3. **Algorithm Design**: Fibonacci heaps, search algorithms
4. **Art and Architecture**: Golden ratio proportions

### Knapsack Applications
1. **Resource Allocation**: CPU scheduling, memory management
2. **Financial Planning**: Portfolio optimization, budget allocation
3. **Cutting Stock Problem**: Material usage optimization
4. **Cargo Loading**: Maximizing value while respecting weight constraints

### N-Queens Applications
1. **Constraint Satisfaction**: General CSP solving techniques
2. **Game AI**: Board game state evaluation
3. **Scheduling**: Task assignment with conflict constraints
4. **Circuit Design**: Component placement optimization

### Permutation Applications
1. **Cryptography**: Encryption key generation
2. **Optimization**: Traveling salesman problem variants
3. **Testing**: Generating test cases for all input combinations
4. **Bioinformatics**: DNA sequence analysis

## Optimization Techniques

### Memoization Best Practices
```go
type MemoTable struct {
    cache map[string]interface{}
    mutex sync.RWMutex
}

func (m *MemoTable) Get(key string) (interface{}, bool) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    val, exists := m.cache[key]
    return val, exists
}

func (m *MemoTable) Set(key string, value interface{}) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.cache[key] = value
}
```

### Tail Recursion Optimization
```go
func factorialTailRec(n, acc int) int {
    if n <= 1 {
        return acc
    }
    return factorialTailRec(n-1, n*acc)
}
```

### Iterative Conversion
```go
func fibonacciIterative(n int) int {
    if n <= 1 {
        return n
    }
    
    prev, curr := 0, 1
    for i := 2; i <= n; i++ {
        prev, curr = curr, prev+curr
    }
    return curr
}
```

### Stack Overflow Prevention
```go
func factorialSafe(n int) (int, error) {
    const maxN = 1000 // Prevent stack overflow
    if n > maxN {
        return 0, fmt.Errorf("input too large: %d > %d", n, maxN)
    }
    return factorialImpl(n), nil
}
```

## Common Pitfalls and Error Handling

### 1. Stack Overflow
```go
func fibonacci(n int) int {
    if n > 1000 {
        panic("input too large, risk of stack overflow")
    }
    // ... implementation
}
```

### 2. Integer Overflow
```go
func factorial(n int) int {
    if n > 20 { // 21! > max int64
        return -1 // or use big.Int
    }
    // ... implementation
}
```

### 3. Invalid Input Handling
```go
func knapsack(weights, values []int, capacity int) (int, error) {
    if len(weights) != len(values) {
        return 0, errors.New("weights and values must have same length")
    }
    if capacity < 0 {
        return 0, errors.New("capacity cannot be negative")
    }
    // ... implementation
}
```

### 4. Memory Management
```go
func permuteLarge(nums []int) [][]int {
    if len(nums) > 10 {
        panic("too many permutations, would exceed memory")
    }
    // ... implementation
}
```

## Conclusion

Recursive algorithms provide elegant solutions to complex problems by breaking them down into smaller, similar subproblems. Key insights:

1. **Problem Decomposition**: Recursion naturally handles problems with recursive structure
2. **Optimization**: Memoization and tabulation can dramatically improve performance
3. **Trade-offs**: Consider memory usage, stack depth, and performance characteristics
4. **Real-world Applicability**: Many algorithms have direct applications in various domains

Understanding recursion is fundamental to algorithm design, dynamic programming, and problem-solving in computer science. The examples demonstrate both the power and the importance of careful optimization in recursive solutions.