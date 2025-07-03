package helper

func setBit(n int, pos uint) int {
	return n | (1 << pos)
}
