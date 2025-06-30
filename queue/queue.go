package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	return item, true
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) size() int {
	return len(q.items)
}
