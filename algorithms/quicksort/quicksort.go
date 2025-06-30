package quicksort

func quicksort(arr []int, low int, high int) {
	if low < high {
		// Partition array
		pivotIndex := partition(arr, low, high)
		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high] // Choose the last element as pivot
	i := low - 1       // Pointer for the smaller element

	for j := low; j < high; j++ {
		if arr[j] < pivot { // If current element is smaller than pivot
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap pivot to the correct position
	return i + 1                              // Return the partition index
}
