package circularQueue

type CircularQueue struct {
	items []int
	front int
	rear  int
	size  int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, capacity),
		front: 0,
		rear:  -1,
		size:  0,
	}
}

func (cq *CircularQueue) Enqueue(item int) bool {
	if cq.size == len(cq.items) {
		return false // Queue is full
	}
	cq.rear = (cq.rear + 1) % len(cq.items)
	cq.items[cq.rear] = item
	cq.size++
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.size == 0 {
		return 0, false // Queue is empty
	}
	item := cq.items[cq.front]
	cq.front = (cq.front + 1) % len(cq.items)
	cq.size--
	return item, true
}

func (cq *CircularQueue) Front() (int, bool) {
	if cq.size == 0 {
		return 0, false // Queue is empty
	}
	return cq.items[cq.front], true
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}

func (cq *CircularQueue) IsFull() bool {
	return cq.size == len(cq.items)
}

func (cq *CircularQueue) Size() int {
	return cq.size
}
