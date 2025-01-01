package heap

import "errors"

// Item represents an element in the priority queue with a value and priority
type Item struct {
	Value    interface{}
	Priority int
	Index    int // Insertion order for FIFO behavior
}

// PriorityQueue implements a priority queue using a min heap
type PriorityQueue struct {
	items []*Item
	size  int // Track size for insertion order
}

// NewPriorityQueue creates and returns a new instance of PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items: make([]*Item, 0),
		size:  0,
	}
}

// Enqueue adds a new item to the priority queue
func (pq *PriorityQueue) Enqueue(value interface{}, priority int) {
	item := &Item{
		Value:    value,
		Priority: priority,
		Index:    pq.size,
	}
	pq.size++
	pq.items = append(pq.items, item)
	pq.heapifyUp(len(pq.items) - 1)
}

// Dequeue removes and returns the highest priority item
func (pq *PriorityQueue) Dequeue() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue is empty")
	}

	item := pq.items[0]
	lastIdx := len(pq.items) - 1
	pq.items[0] = pq.items[lastIdx]
	pq.items = pq.items[:lastIdx]

	if !pq.IsEmpty() {
		pq.heapifyDown(0)
	}

	return item.Value, nil
}

// Peek returns the highest priority item without removing it
func (pq *PriorityQueue) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue is empty")
	}
	return pq.items[0].Value, nil
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue) Size() int {
	return len(pq.items)
}

// IsEmpty returns true if the priority queue contains no elements
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

// heapifyUp maintains the min-heap property by moving a node up the tree
func (pq *PriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parentIdx := parent(index)
		if pq.items[parentIdx].Priority > pq.items[index].Priority ||
			(pq.items[parentIdx].Priority == pq.items[index].Priority &&
				pq.items[parentIdx].Index > pq.items[index].Index) {
			pq.items[parentIdx], pq.items[index] = pq.items[index], pq.items[parentIdx]
			index = parentIdx
		} else {
			break
		}
	}
}

// heapifyDown maintains the min-heap property by moving a node down the tree
func (pq *PriorityQueue) heapifyDown(index int) {
	lastIdx := len(pq.items) - 1
	for {
		smallest := index
		leftIdx := leftChild(index)
		rightIdx := rightChild(index)

		if leftIdx <= lastIdx && (pq.items[leftIdx].Priority < pq.items[smallest].Priority ||
			(pq.items[leftIdx].Priority == pq.items[smallest].Priority &&
				pq.items[leftIdx].Index < pq.items[smallest].Index)) {
			smallest = leftIdx
		}
		if rightIdx <= lastIdx && (pq.items[rightIdx].Priority < pq.items[smallest].Priority ||
			(pq.items[rightIdx].Priority == pq.items[smallest].Priority &&
				pq.items[rightIdx].Index < pq.items[smallest].Index)) {
			smallest = rightIdx
		}

		if smallest == index {
			break
		}

		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		index = smallest
	}
}
