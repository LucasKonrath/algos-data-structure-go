package slidingwindow

import (
	"testing"
)

func TestMaxSumSubArray_NormalCase(t *testing.T) {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	expected := 9 // 5+1+3
	result := maxSumSubArray(arr, k)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxSumSubArray_KEqualsArrayLength(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	k := 4
	expected := 10
	result := maxSumSubArray(arr, k)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxSumSubArray_KGreaterThanArrayLength(t *testing.T) {
	arr := []int{1, 2}
	k := 3
	expected := 0
	result := maxSumSubArray(arr, k)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxSumSubArray_EmptyArray(t *testing.T) {
	arr := []int{}
	k := 1
	expected := 0
	result := maxSumSubArray(arr, k)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
