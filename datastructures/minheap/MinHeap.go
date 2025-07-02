package minheap

type MinHeap struct {
	items []int
}

func (h *MinHeap) ParentIndex(index int) int {
	if index == 0 {
		return -1 // No parent for the root
	}
	return (index - 1) / 2
}

func (h *MinHeap) LeftChildIndex(index int) int {
	return 2*index + 1
}

func (h *MinHeap) RightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MinHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinHeap) Insert(value int) {
	h.items = append(h.items, value)
	if len(h.items) > 1 {
		h.heapifyUp(len(h.items) - 1)
	}
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.items[h.ParentIndex(index)] > h.items[index] {
		h.Swap(h.ParentIndex(index), index)
		index = h.ParentIndex(index)
	}
}

func (h *MinHeap) Extract() int {
	if len(h.items) == 0 {
		return 0 // or error value
	}
	extracted := h.items[0]
	lastIndex := len(h.items) - 1
	if lastIndex == 0 {
		h.items = nil
		return extracted
	}
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex] // Remove the last element
	h.heapifyDown(0)
	return extracted
}

func (h *MinHeap) Size() int {
	return len(h.items)
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.items) - 1
	l, r := h.LeftChildIndex(index), h.RightChildIndex(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex { // Only left child exists
			childToCompare = l
		} else if h.items[l] < h.items[r] { // Compare left and right children
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.items[index] > h.items[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = h.LeftChildIndex(index), h.RightChildIndex(index)
		} else {
			return
		}
	}
}
