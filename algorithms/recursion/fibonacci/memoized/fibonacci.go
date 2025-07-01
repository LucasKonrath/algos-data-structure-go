package memoized

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

func fibonacciWrapper(n int) int {
	memo := make(map[int]int)
	return fibonacci(n, memo)
}
