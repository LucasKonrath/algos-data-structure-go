package tabulation

func fibonacciTabulation(n int) int {
	if n <= 1 {
		return n
	}

	// Create a table to store Fibonacci numbers up to n
	fibTable := make([]int, n+1)
	fibTable[0], fibTable[1] = 0, 1

	// Fill the table in a bottom-up manner
	for i := 2; i <= n; i++ {
		fibTable[i] = fibTable[i-1] + fibTable[i-2]
	}

	return fibTable[n]
}
