package slidingwindow

func smallestSubArrayWithSum(arr []int, target int) int {
	start := 0
	end := 0
	sum := 0
	minLength := len(arr) + 1 // Initialize to a value larger than any possible subarray length

	for end < len(arr) {
		for sum < target && end < len(arr) {
			sum += arr[end]
			end++
		}

		for sum >= target && start < end {
			if end-start < minLength {
				minLength = end - start
			}
			sum -= arr[start]
			start++
		}
	}
	if minLength > len(arr) {
		return 0 // If no valid subarray found, return 0
	}
	return minLength
}
