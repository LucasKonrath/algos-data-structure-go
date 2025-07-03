package helper

func toggleBit(n int, pos uint) int {
	return n ^ (1 << pos) // XOR operation toggles the bit at position pos
}
