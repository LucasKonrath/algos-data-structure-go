package optimization

// Go allows two return values, which is useful for functions that need to return multiple results.
func minMax(arr []int) (min, max int) {
	if len(arr) == 0 {
		return 0, 0 // Handle empty array case
	}

	min = arr[0]
	max = arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
