package bitmanipulation

func isPoweOfTwo(n int) bool {
	// A number is a power of two if it is greater than 0 and has only one bit set
	return n > 0 && (n&(n-1)) == 0 // n & (n-1) will be 0 if n is a power of two
}
