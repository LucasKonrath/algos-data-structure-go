package deque

type Deque struct {
	items []int
}

func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) PopFront() (int, bool) {
	if len(d.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

func (d *Deque) PopBack() (int, bool) {
	if len(d.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

func (d *Deque) Front() (int, bool) {
	if len(d.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	return d.items[0], true
}

func (d *Deque) Back() (int, bool) {
	if len(d.items) == 0 {
		return 0, false // or handle underflow as needed
	}
	return d.items[len(d.items)-1], true
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}
