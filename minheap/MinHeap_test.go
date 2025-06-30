package minheap

import "testing"

func TestMinHeap_InsertAndExtract(t *testing.T) {
	h := &MinHeap{}
	values := []int{5, 3, 8, 1, 2, 7}
	for _, v := range values {
		h.Insert(v)
	}

	expectedOrder := []int{1, 2, 3, 5, 7, 8}
	for _, expected := range expectedOrder {
		extracted := h.Extract()
		if extracted != expected {
			t.Errorf("Expected %d, got %d", expected, extracted)
		}
	}

}

func TestMinHeap_ExtractFromEmpty(t *testing.T) {
	h := &MinHeap{}
	extracted := h.Extract()
	if extracted != 0 {
		t.Errorf("Expected 0 when extracting from empty heap, got %d", extracted)
	}
}
