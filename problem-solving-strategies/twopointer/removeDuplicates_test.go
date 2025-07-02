package twopointer

import "testing"

func TestRemoveDuplicates_Empty(t *testing.T) {
	arr := []int{}
	expected := 0
	result := removeDuplicates(arr)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestRemoveDuplicates_SingleElement(t *testing.T) {
	arr := []int{1}
	expected := 1
	result := removeDuplicates(arr)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestRemoveDuplicates_NoDuplicates(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := 4
	result := removeDuplicates(arr)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	for i := 0; i < result; i++ {
		if arr[i] != i+1 {
			t.Errorf("Expected arr[%d]=%d, got %d", i, i+1, arr[i])
		}
	}
}

func TestRemoveDuplicates_WithDuplicates(t *testing.T) {
	arr := []int{1, 1, 2, 2, 3, 3, 3, 4}
	expected := 4
	result := removeDuplicates(arr)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	expectedArr := []int{1, 2, 3, 4}
	for i := 0; i < result; i++ {
		if arr[i] != expectedArr[i] {
			t.Errorf("Expected arr[%d]=%d, got %d", i, expectedArr[i], arr[i])
		}
	}
}

func TestRemoveDuplicates_AllSame(t *testing.T) {
	arr := []int{5, 5, 5, 5, 5}
	expected := 1
	result := removeDuplicates(arr)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
	if arr[0] != 5 {
		t.Errorf("Expected arr[0]=5, got %d", arr[0])
	}
}
