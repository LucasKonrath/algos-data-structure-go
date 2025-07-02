package slidingwindow

func maxSumSubArray(arr []int, k int) int {
	if len(arr) < k {
		return 0 // Not enough elements
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i] // Calculate the sum of the first 'k' elements
	}
	maxSum := windowSum

	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i] // Slide the window by removing the first element and adding the next element
		if windowSum > maxSum {
			maxSum = windowSum // Update maxSum if the current window sum is greater
		}
	}
	return maxSum
}
