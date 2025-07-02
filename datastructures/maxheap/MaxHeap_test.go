package maxheap

import "testing"

func TestMaxHeap_InsertAndExtract(t *testing.T) {
	heap := &MaxHeap{}
	values := []int{5, 3, 8, 1, 2, 7}
	for _, v := range values {
		heap.Insert(v)
	}

	expectedOrder := []int{8, 7, 5, 3, 2, 1}
	for _, expected := range expectedOrder {
		extracted := heap.Extract()
		if extracted != expected {
			t.Errorf("Expected %d, got %d", expected, extracted)
		}
	}

}

func TestMaxHeap_ExtractFromEmpty(t *testing.T) {
	heap := &MaxHeap{}
	extracted := heap.Extract()
	if extracted != 0 {
		t.Errorf("Expected 0 when extracting from empty heap, got %d", extracted)
	}
}
