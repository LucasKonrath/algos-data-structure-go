package slidingwindow

import "testing"

func TestSmallestSubarrayWithSum_NormalCase(t *testing.T) {
	arr := []int{2, 1, 5, 2, 3, 2}
	target := 7
	expected := 2 // [5,2] or [2,5]
	result := smallestSubArrayWithSum(arr, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSmallestSubarrayWithSum_NoValidSubarray(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	target := 10
	expected := 0
	result := smallestSubArrayWithSum(arr, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSmallestSubarrayWithSum_WholeArrayNeeded(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	target := 10
	expected := 4
	result := smallestSubArrayWithSum(arr, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSmallestSubarrayWithSum_SingleElement(t *testing.T) {
	arr := []int{8, 1, 2, 3}
	target := 8
	expected := 1
	result := smallestSubArrayWithSum(arr, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSmallestSubarrayWithSum_EmptyArray(t *testing.T) {
	arr := []int{}
	target := 5
	expected := 0
	result := smallestSubArrayWithSum(arr, target)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
