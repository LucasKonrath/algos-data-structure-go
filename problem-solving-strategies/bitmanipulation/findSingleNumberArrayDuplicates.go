package bitmanipulation

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR operation
	}
	// The XOR operation will cancel out all numbers that appear twice,
	// leaving only the number that appears once.
	// This works because:
	// - a ^ a = 0 (any number XORed with itself is 0)
	// - a ^ 0 = a (any number XORed with 0 is the number itself)
	// - XOR is commutative and associative, so the order of operations doesn't matter.
	// Example: For nums = [4, 1, 2, 1, 2], the operations would be:
	// 0 ^ 4 = 4
	// 4 ^ 1 = 5
	// 5 ^ 2 = 7
	// 7 ^ 1 = 6
	// 6 ^ 2 = 4
	// Result will be 4, which is the single number.
	return result
}
