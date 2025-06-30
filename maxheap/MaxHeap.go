package maxheap

type MaxHeap struct {
	items []int
}

func (mh *MaxHeap) ParentIndex(index int) int {
	if index == 0 {
		return -1 // No parent for the root
	}
	return (index - 1) / 2
}

func (mh *MaxHeap) LeftChildIndex(index int) int {
	return 2*index + 1
}

func (mh *MaxHeap) RightChildIndex(index int) int {
	return 2*index + 2
}

func (mh *MaxHeap) Swap(i, j int) {
	mh.items[i], mh.items[j] = mh.items[j], mh.items[i]
}

func (mh *MaxHeap) Insert(value int) {
	mh.items = append(mh.items, value)
	if len(mh.items) > 1 {
		mh.bubbleUp(len(mh.items) - 1)
	}
}

func (mh *MaxHeap) bubbleUp(index int) {
	for index > 0 && mh.items[mh.ParentIndex(index)] < mh.items[index] {
		mh.Swap(mh.ParentIndex(index), index)
		index = mh.ParentIndex(index)
	}
}

func (mh *MaxHeap) Extract() int {
	if len(mh.items) == 0 {
		return 0 // or error value
	}
	extracted := mh.items[0]
	lastIndex := len(mh.items) - 1
	if lastIndex == 0 {
		mh.items = nil
		return extracted
	}
	mh.items[0] = mh.items[lastIndex]
	mh.items = mh.items[:lastIndex] // Remove the last element
	mh.bubbleDown(0)
	return extracted
}

func (mh *MaxHeap) bubbleDown(index int) {
	for {
		left := mh.LeftChildIndex(index)
		right := mh.RightChildIndex(index)
		largest := index

		if left < len(mh.items) && mh.items[left] > mh.items[largest] {
			largest = left
		}
		if right < len(mh.items) && mh.items[right] > mh.items[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		mh.Swap(index, largest)
		index = largest
	}
}

func (mh *MaxHeap) Size() int {
	return len(mh.items)
}
