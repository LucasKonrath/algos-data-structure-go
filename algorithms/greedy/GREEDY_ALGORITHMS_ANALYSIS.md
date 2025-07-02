# Greedy Algorithms - Comprehensive Analysis

## Table of Contents
1. [Algorithm Overview](#algorithm-overview)
2. [Coin Change Algorithm](#coin-change-algorithm)
3. [Fractional Knapsack Algorithm](#fractional-knapsack-algorithm)
4. [Mathematical Foundations](#mathematical-foundations)
5. [Performance Comparison](#performance-comparison)
6. [Real-World Applications](#real-world-applications)
7. [Common Pitfalls and Limitations](#common-pitfalls-and-limitations)

## Algorithm Overview

### Greedy Algorithm Principles
Greedy algorithms make locally optimal choices at each step with the hope of finding a global optimum. They follow the greedy choice property and optimal substructure:

1. **Greedy Choice Property**: A global optimum can be arrived at by making a locally optimal choice
2. **Optimal Substructure**: An optimal solution contains optimal solutions to subproblems
3. **Irrevocability**: Once a choice is made, it cannot be undone

### Key Characteristics
- **Simple Implementation**: Generally easier to implement than dynamic programming
- **Efficiency**: Often run in linear or polynomial time
- **No Backtracking**: Never reconsider previous choices
- **Problem-Specific**: Not all problems have greedy solutions

## Coin Change Algorithm

### Implementation Analysis

```go
func coinchange(coins []int, amount int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Sort coins in descending order
    count := 0
    for _, coin := range coins {
        for amount >= coin {
            amount -= coin
            count++
        }
    }
    if amount == 0 {
        return count
    }
    return -1
}
```

### Algorithm Breakdown

#### Step 1: Sorting
- **Time Complexity**: O(n log n) where n is the number of coin denominations
- **Purpose**: Ensures we try larger denominations first for optimal greedy choice

#### Step 2: Greedy Selection
- **Time Complexity**: O(amount / smallest_coin) in worst case
- **Strategy**: Use as many of the largest denomination as possible, then move to smaller denominations

#### Step 3: Validation
- **Purpose**: Check if exact change is possible
- **Return**: Number of coins used or -1 if impossible

### Complexity Analysis
- **Time Complexity**: O(n log n + amount)
- **Space Complexity**: O(1) - only using constant extra space
- **Optimality**: Only optimal for canonical coin systems (like standard currency)

### Execution Trace Example
**Input**: coins = [1, 2, 5], amount = 11

```
Initial: coins = [5, 2, 1] (sorted), amount = 11, count = 0

Step 1: coin = 5
  - amount >= 5? Yes (11 >= 5)
  - amount = 11 - 5 = 6, count = 1
  - amount >= 5? Yes (6 >= 5)  
  - amount = 6 - 5 = 1, count = 2
  - amount >= 5? No (1 < 5)

Step 2: coin = 2
  - amount >= 2? No (1 < 2)

Step 3: coin = 1
  - amount >= 1? Yes (1 >= 1)
  - amount = 1 - 1 = 0, count = 3
  - amount >= 1? No (0 < 1)

Final: amount = 0, return count = 3
Result: [5, 5, 1] using 3 coins
```

### Limitations
This greedy approach doesn't work for all coin systems. For example:
- **Coins**: [1, 3, 4], **Amount**: 6
- **Greedy Result**: [4, 1, 1] = 3 coins
- **Optimal Result**: [3, 3] = 2 coins

## Fractional Knapsack Algorithm

### Implementation Analysis

```go
type Item struct {
    value, weight float64
}

func fractionalKnapsack(capacity float64, items []Item) float64 {
    // Sort items by value-to-weight ratio in descending order
    sort.Slice(items, func(i, j int) bool {
        return (items[i].value / items[i].weight) > (items[j].value / items[j].weight)
    })

    totalValue := 0.0
    for _, item := range items {
        if capacity <= 0 {
            break
        }
        if item.weight <= capacity {
            totalValue += item.value
            capacity -= item.weight
        } else {
            totalValue += item.value * (capacity / item.weight)
            capacity = 0
        }
    }
    return totalValue
}
```

### Algorithm Breakdown

#### Step 1: Calculate Value-to-Weight Ratios
Each item's efficiency is measured by its value per unit weight.

#### Step 2: Sort by Efficiency
- **Time Complexity**: O(n log n)
- **Purpose**: Prioritize most valuable items per unit weight

#### Step 3: Greedy Selection
- **Full Item**: Take entire item if it fits
- **Fractional Item**: Take fraction that fills remaining capacity
- **Time Complexity**: O(n)

### Complexity Analysis
- **Time Complexity**: O(n log n) - dominated by sorting
- **Space Complexity**: O(1) - sorting in place
- **Optimality**: Always optimal for fractional knapsack

### Execution Trace Example
**Input**: capacity = 50, items = [{60, 10}, {100, 20}, {120, 30}]

```
Step 1: Calculate ratios
  - Item 1: 60/10 = 6.0
  - Item 2: 100/20 = 5.0  
  - Item 3: 120/30 = 4.0

Step 2: Sort by ratio (descending)
  - Items: [{60, 10}, {100, 20}, {120, 30}]
  - Already sorted in this case

Step 3: Fill knapsack
  - capacity = 50, totalValue = 0

  Item 1 (value=60, weight=10):
    - weight <= capacity? Yes (10 <= 50)
    - totalValue = 0 + 60 = 60
    - capacity = 50 - 10 = 40

  Item 2 (value=100, weight=20):
    - weight <= capacity? Yes (20 <= 40)
    - totalValue = 60 + 100 = 160
    - capacity = 40 - 20 = 20

  Item 3 (value=120, weight=30):
    - weight <= capacity? No (30 > 20)
    - Take fraction: 20/30 = 2/3
    - totalValue = 160 + 120 * (2/3) = 160 + 80 = 240
    - capacity = 0

Final: totalValue = 240.0
```

### Why Greedy Works Here
The fractional knapsack has the greedy choice property because:
1. We can always improve by swapping less efficient items for more efficient ones
2. Taking fractions allows us to fully utilize capacity
3. The optimal solution must prioritize highest value-to-weight ratios

## Mathematical Foundations

### Greedy Choice Property Proof (Fractional Knapsack)

**Theorem**: The greedy algorithm produces an optimal solution for fractional knapsack.

**Proof Sketch**:
1. Let G be the greedy solution and O be any optimal solution
2. If G ≠ O, there exists a first item where they differ
3. In O, we can replace less efficient items with more efficient ones without decreasing total value
4. This process converges to the greedy solution, proving optimality

### Exchange Argument
For any non-greedy optimal solution:
- We can "exchange" portions of less efficient items for more efficient ones
- This exchange never decreases the total value
- Repeated exchanges lead to the greedy solution

## Performance Comparison

### Time Complexity Summary
| Algorithm | Sorting | Selection | Total |
|-----------|---------|-----------|-------|
| Coin Change | O(n log n) | O(amount) | O(n log n + amount) |
| Fractional Knapsack | O(n log n) | O(n) | O(n log n) |

### Space Complexity
- **Coin Change**: O(1) - constant space
- **Fractional Knapsack**: O(1) - sorting in place

### Comparison with Alternative Approaches

#### Coin Change vs Dynamic Programming
- **Greedy**: O(n log n + amount) time, O(1) space
- **DP**: O(n × amount) time, O(amount) space
- **Trade-off**: Greedy is faster but only works for canonical coin systems

#### Fractional vs 0/1 Knapsack  
- **Fractional**: O(n log n) time, always optimal with greedy
- **0/1 Knapsack**: O(n × capacity) time with DP, greedy doesn't work

## Real-World Applications

### Coin Change Algorithm
1. **Currency Exchange Systems**
   - ATM cash dispensing
   - Point-of-sale change calculation
   - Vending machine operations

2. **Resource Allocation**
   - Bandwidth allocation with fixed units
   - Memory page allocation
   - CPU time slice distribution

3. **Manufacturing**
   - Material cutting optimization
   - Batch size optimization
   - Inventory management

### Fractional Knapsack Algorithm
1. **Financial Portfolio Management**
   - Asset allocation with continuous investments
   - Risk-return optimization
   - Commodity trading

2. **Resource Management**
   - Network bandwidth allocation
   - CPU scheduling with preemption
   - Memory allocation systems

3. **Supply Chain Optimization**
   - Truck loading with divisible goods
   - Warehouse space allocation
   - Production planning

4. **Chemical Engineering**
   - Mixture optimization
   - Reaction yield maximization
   - Process optimization

## Common Pitfalls and Limitations

### Coin Change Pitfalls
1. **Non-Canonical Coin Systems**
   ```go
   // This will fail with greedy approach
   coins := []int{1, 3, 4}
   amount := 6
   // Greedy: [4, 1, 1] = 3 coins
   // Optimal: [3, 3] = 2 coins
   ```

2. **Integer Overflow**
   - Large amounts may cause overflow
   - Need to handle edge cases properly

3. **Empty Coin Set**
   - Must validate input parameters
   - Handle zero amount correctly

### Fractional Knapsack Limitations
1. **Indivisible Items**
   - Real-world items often cannot be divided
   - May need 0/1 knapsack approach instead

2. **Setup Costs**
   - Algorithm assumes no setup cost for taking fractions
   - Real systems may have transaction costs

3. **Floating Point Precision**
   - Accumulation of rounding errors
   - Need careful comparison with epsilon values

### General Greedy Algorithm Warnings
1. **Problem Suitability**
   - Not all optimization problems have greedy solutions
   - Must prove greedy choice property and optimal substructure

2. **Local vs Global Optimum**
   - Greedy makes locally optimal choices
   - May miss globally optimal solutions

3. **Implementation Details**
   - Sorting stability may matter
   - Edge case handling is crucial

### Error Handling Best Practices
```go
func coinChangeRobust(coins []int, amount int) (int, error) {
    if amount < 0 {
        return 0, errors.New("amount cannot be negative")
    }
    if amount == 0 {
        return 0, nil
    }
    if len(coins) == 0 {
        return -1, errors.New("no coins available")
    }
    
    // Filter out invalid coins
    validCoins := make([]int, 0, len(coins))
    for _, coin := range coins {
        if coin > 0 {
            validCoins = append(validCoins, coin)
        }
    }
    
    if len(validCoins) == 0 {
        return -1, errors.New("no valid coins")
    }
    
    // Proceed with algorithm...
}
```

## Conclusion

Greedy algorithms provide efficient solutions for specific problem classes but require careful analysis to ensure optimality. The coin change and fractional knapsack problems demonstrate both the power and limitations of the greedy approach:

- **Strengths**: Simple, efficient, intuitive
- **Weaknesses**: Limited applicability, requires proof of optimality
- **Key Insight**: Success depends on problem structure, not algorithm sophistication

Understanding when and why greedy algorithms work is crucial for algorithm design and optimization problem solving.