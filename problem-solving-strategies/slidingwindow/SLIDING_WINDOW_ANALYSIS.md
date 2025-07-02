# Sliding Window Technique - Expert Analysis

## Technique Overview
The **Sliding Window** technique is a computational pattern that transforms nested loops into a single loop, reducing time complexity from O(n²) or O(n³) to O(n) for problems involving contiguous subarrays or substrings.

## Core Principle
Maintain a "window" of elements and slide it across the data structure, dynamically adjusting the window size based on problem constraints while avoiding redundant calculations.

## Window Types

### 1. Fixed-Size Window
```go
func maxSumSubarray(arr []int, k int) int {
    windowSum := 0
    // Calculate sum of first window
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    
    maxSum := windowSum
    // Slide the window
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        maxSum = max(maxSum, windowSum)
    }
    return maxSum
}
```

### 2. Variable-Size Window (Expanding/Contracting)
```go
func longestSubstringWithoutRepeats(s string) int {
    charMap := make(map[rune]int)
    left, maxLen := 0, 0
    
    for right, char := range s {
        if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
            left = lastIndex + 1  // Contract window
        }
        charMap[char] = right
        maxLen = max(maxLen, right-left+1)  // Expand window
    }
    return maxLen
}
```

## Implementation Analysis: Longest Unique Substring

### Algorithm Breakdown
```go
func longestUniquesSubstring(s string) string {
    charIndex := make(map[rune]int)  // Character → Last seen index
    start := 0                       // Window left boundary
    maxLength := 0                   // Maximum window size found
    maxStart := 0                    // Start of optimal window
    
    for i, char := range s {
        // Window contraction condition
        if lastIndex, found := charIndex[char]; found && lastIndex >= start {
            start = lastIndex + 1
        }
        
        // Window expansion and tracking
        if i-start+1 > maxLength {
            maxLength = i - start + 1
            maxStart = start
        }
        
        charIndex[char] = i  // Update character position
    }
    return s[maxStart : maxStart+maxLength]
}
```

### Execution Trace Example
```
Input: "abcabcbb"

i=0, char='a': charIndex={a:0}, start=0, window="a", maxLen=1
i=1, char='b': charIndex={a:0,b:1}, start=0, window="ab", maxLen=2  
i=2, char='c': charIndex={a:0,b:1,c:2}, start=0, window="abc", maxLen=3
i=3, char='a': found 'a' at 0 >= start(0), start=1, window="bca", maxLen=3
i=4, char='b': found 'b' at 1 >= start(1), start=2, window="cab", maxLen=3
i=5, char='c': found 'c' at 2 >= start(2), start=3, window="abc", maxLen=3
i=6, char='b': charIndex={a:3,b:6,c:5}, start=3, window="abcb", maxLen=4
i=7, char='b': found 'b' at 6 >= start(3), start=7, window="b", maxLen=4

Result: "abcb" (or any 4-character unique substring)
```

## Complexity Analysis

### Time Complexity
```
Single Pass: O(n)
- Each element visited at most twice (once by right pointer, once by left pointer)
- Hash map operations: O(1) average case
- String slicing: O(k) where k is result length
Total: O(n)
```

### Space Complexity
```
Hash Map Storage: O(min(m, n))
- m = size of character set (e.g., 256 for ASCII)
- n = length of string
- Map stores at most min(m, n) unique characters
```

## Window State Management

### Window Invariants
1. **Uniqueness**: No duplicate characters in current window
2. **Maximality**: Window cannot be extended without violating uniqueness
3. **Optimality**: Longest valid window ending at current position

### State Transitions
```
State: [start, i] represents current window

Transitions:
1. Extend: i++ (always happens)
2. Contract: start = lastIndex[char] + 1 (when duplicate found)
3. Update: Record new maximum if current window is larger
```

## Pattern Variations

### Template 1: Fixed Window Size
```go
func slidingWindowFixed(arr []int, k int) []int {
    result := []int{}
    windowSum := 0
    
    // Initialize first window
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    result = append(result, windowSum)
    
    // Slide window
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        result = append(result, windowSum)
    }
    return result
}
```

### Template 2: Variable Window (Condition-based)
```go
func slidingWindowVariable(arr []int, condition func([]int) bool) int {
    left, maxLen := 0, 0
    
    for right := 0; right < len(arr); right++ {
        // Expand window by including arr[right]
        
        // Contract window while condition is violated
        for !condition(arr[left:right+1]) {
            left++
        }
        
        // Update maximum window size
        maxLen = max(maxLen, right-left+1)
    }
    return maxLen
}
```

### Template 3: Two-Pointer Technique
```go
func twoPointerSlidingWindow(arr []int, target int) bool {
    left, right := 0, 0
    currentSum := 0
    
    for right < len(arr) {
        currentSum += arr[right]
        
        // Shrink window if needed
        for currentSum > target && left <= right {
            currentSum -= arr[left]
            left++
        }
        
        if currentSum == target {
            return true
        }
        right++
    }
    return false
}
```

## Advanced Applications

### Minimum Window Substring
```go
func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }
    
    // Frequency map for target string
    tFreq := make(map[rune]int)
    for _, char := range t {
        tFreq[char]++
    }
    
    required := len(tFreq)  // Unique characters in t
    formed := 0             // Characters with desired frequency in window
    
    windowCounts := make(map[rune]int)
    left, right := 0, 0
    
    // (window length, left, right)
    ans := []int{math.MaxInt, 0, 0}
    
    for right < len(s) {
        char := rune(s[right])
        windowCounts[char]++
        
        // Check if frequency matches desired count
        if count, exists := tFreq[char]; exists && windowCounts[char] == count {
            formed++
        }
        
        // Contract window if all characters are matched
        for left <= right && formed == required {
            char = rune(s[left])
            
            // Update result if this window is smaller
            if right-left+1 < ans[0] {
                ans[0] = right - left + 1
                ans[1] = left
                ans[2] = right
            }
            
            windowCounts[char]--
            if count, exists := tFreq[char]; exists && windowCounts[char] < count {
                formed--
            }
            left++
        }
        right++
    }
    
    if ans[0] == math.MaxInt {
        return ""
    }
    return s[ans[1]:ans[2]+1]
}
```

### Sliding Window Maximum
```go
import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
    deque := list.New()  // Stores indices
    result := []int{}
    
    for i, num := range nums {
        // Remove indices outside current window
        for deque.Len() > 0 && deque.Front().Value.(int) < i-k+1 {
            deque.Remove(deque.Front())
        }
        
        // Remove indices with smaller values (they'll never be maximum)
        for deque.Len() > 0 && nums[deque.Back().Value.(int)] < num {
            deque.Remove(deque.Back())
        }
        
        deque.PushBack(i)
        
        // Add maximum to result when window is fully formed
        if i >= k-1 {
            result = append(result, nums[deque.Front().Value.(int)])
        }
    }
    return result
}
```

## Performance Optimization

### Memory-Efficient Character Tracking
```go
func longestUniqueSubstringOptimized(s string) int {
    // Use array instead of map for ASCII characters
    charIndex := make([]int, 256)
    for i := range charIndex {
        charIndex[i] = -1  // Initialize with invalid index
    }
    
    left, maxLen := 0, 0
    
    for right := 0; right < len(s); right++ {
        char := s[right]
        
        if charIndex[char] >= left {
            left = charIndex[char] + 1
        }
        
        charIndex[char] = right
        maxLen = max(maxLen, right-left+1)
    }
    return maxLen
}
```

### Bit Manipulation for Small Character Sets
```go
func hasUniqueCharacters(s string) bool {
    var bitVector int
    
    for _, char := range s {
        bit := 1 << (char - 'a')
        if bitVector&bit != 0 {
            return false  // Duplicate found
        }
        bitVector |= bit
    }
    return true
}
```

## Common Pitfalls and Solutions

### 1. Index Management
```go
// Wrong: Off-by-one errors
windowSize := right - left  // Missing +1

// Correct: Inclusive window size
windowSize := right - left + 1
```

### 2. Hash Map Key Types
```go
// Wrong: Comparing different types
charMap := make(map[byte]int)  // byte keys
for i, char := range s {       // char is rune
    if _, exists := charMap[char]; exists {  // Type mismatch
        
// Correct: Consistent types
charMap := make(map[rune]int)
for i, char := range s {
    if _, exists := charMap[char]; exists {
```

### 3. Window Contraction Logic
```go
// Wrong: Moving left pointer incorrectly
if duplicate {
    left++  // Only moves one step
}

// Correct: Jump to optimal position
if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
    left = lastIndex + 1
}
```

## Mathematical Properties

### Window Count Analysis
For a string of length n:
- **Maximum possible windows**: n(n+1)/2
- **Sliding window visits**: Each position at most twice = 2n
- **Efficiency gain**: From O(n²) to O(n)

### Amortized Analysis
```
Total pointer movements:
- Right pointer: n movements (one per character)
- Left pointer: ≤ n movements (never backtracks)
- Total: ≤ 2n = O(n)
```

## Real-world Applications

### Network Packet Analysis
```go
func detectAnomalousTraffic(packets []int, windowSize int, threshold int) []int {
    anomalies := []int{}
    windowSum := 0
    
    for i, packet := range packets {
        windowSum += packet
        
        if i >= windowSize {
            windowSum -= packets[i-windowSize]
        }
        
        if i >= windowSize-1 && windowSum > threshold {
            anomalies = append(anomalies, i-windowSize+1)
        }
    }
    return anomalies
}
```

### DNA Sequence Analysis
```go
func findRepeatedDNASequences(s string) []string {
    seen := make(map[string]bool)
    repeated := make(map[string]bool)
    
    for i := 0; i <= len(s)-10; i++ {
        sequence := s[i : i+10]
        if seen[sequence] {
            repeated[sequence] = true
        }
        seen[sequence] = true
    }
    
    result := []string{}
    for seq := range repeated {
        result = append(result, seq)
    }
    return result
}
```

### Time Series Analysis
```go
func movingAverage(prices []float64, period int) []float64 {
    if len(prices) < period {
        return []float64{}
    }
    
    averages := make([]float64, len(prices)-period+1)
    windowSum := 0.0
    
    // Initialize first window
    for i := 0; i < period; i++ {
        windowSum += prices[i]
    }
    averages[0] = windowSum / float64(period)
    
    // Slide window
    for i := period; i < len(prices); i++ {
        windowSum = windowSum - prices[i-period] + prices[i]
        averages[i-period+1] = windowSum / float64(period)
    }
    return averages
}
```

## Problem Classification

### By Window Behavior
1. **Fixed-size windows**: Maximum sum subarray, moving average
2. **Expanding windows**: Longest substring with condition
3. **Contracting windows**: Minimum window covering substring
4. **Two-pointer**: Array sum problems, palindrome checking

### By Optimization Target
1. **Maximum**: Longest valid substring/subarray
2. **Minimum**: Shortest valid substring/subarray  
3. **Exact**: Finding specific sum or pattern
4. **Count**: Number of valid windows

The sliding window technique transforms many O(n²) or O(n³) brute force solutions into elegant O(n) algorithms, making it one of the most powerful optimization patterns in competitive programming and real-world applications.