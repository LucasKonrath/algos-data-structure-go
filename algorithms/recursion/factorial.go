package recursion

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
