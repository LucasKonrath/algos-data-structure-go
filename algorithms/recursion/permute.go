package recursion

func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func([]int, []int)

	backtrack = func(current []int, remaining []int) {
		if len(remaining) == 0 {
			result = append(result, append([]int{}, current...)) // Append a copy of current to result
			return
		}

		for i, num := range remaining {
			newCurrent := append(current, num)                                            // Create a new slice with the current number added
			newRemaining := append(append([]int{}, remaining[:i]...), remaining[i+1:]...) // Create a new slice without the current number
			backtrack(newCurrent, newRemaining)                                           // Recur with the new slice
		}
	}

	backtrack([]int{}, nums)
	return result
}
