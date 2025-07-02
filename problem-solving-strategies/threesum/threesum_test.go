package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func sortTriplets(triplets [][]int) {
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

func TestThreeSum_NormalCase(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := threesum(nums)
	sortTriplets(result)
	sortTriplets(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestThreeSum_NoSolution(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := [][]int{}
	result := threesum(nums)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}

func TestThreeSum_AllZeros(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	expected := [][]int{{0, 0, 0}}
	result := threesum(nums)
	sortTriplets(result)
	sortTriplets(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestThreeSum_LessThanThree(t *testing.T) {
	nums := []int{0, 1}
	expected := [][]int{}
	result := threesum(nums)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}

func TestThreeSum_Duplicates(t *testing.T) {
	nums := []int{-2, 0, 0, 2, 2}
	expected := [][]int{{-2, 0, 2}}
	result := threesum(nums)
	sortTriplets(result)
	sortTriplets(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
