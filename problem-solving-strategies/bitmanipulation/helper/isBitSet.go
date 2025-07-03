package helper

func isBitSet(n int, pos uint) bool {
	return (n & (1 << pos)) != 0 // Check if the bit at position pos is set, << operator shifts 1 to the left by pos positions
}
