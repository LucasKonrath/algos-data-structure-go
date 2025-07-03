package bitmanipulation

func clearBit(n int, pos uint) int {
	// Clear the bit at position pos by using bitwise AND with the negation of the mask
	return n &^ (1 << pos)
}
