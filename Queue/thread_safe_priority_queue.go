package queue

import (
	"sync"
)

// PriorityQueueItem represents an item in the priority queue
type PriorityQueueItem[T any] struct {
	Value    T
	Priority int
}

// ThreadSafePriorityQueue represents a thread-safe generic priority queue
type ThreadSafePriorityQueue[T any] struct {
	items []PriorityQueueItem[T]
	size  int
	mutex sync.RWMutex
}

// NewThreadSafePriorityQueue creates a new thread-safe priority queue
func NewThreadSafePriorityQueue[T any]() *ThreadSafePriorityQueue[T] {
	return &ThreadSafePriorityQueue[T]{
		items: make([]PriorityQueueItem[T], 0),
		size:  0,
	}
}

// Enqueue adds an item to the priority queue
func (pq *ThreadSafePriorityQueue[T]) Enqueue(value T, priority int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	item := PriorityQueueItem[T]{Value: value, Priority: priority}
	pq.items = append(pq.items, item)
	pq.size++
	pq.heapifyUp(pq.size - 1)
}

// Dequeue removes and returns the highest priority item
func (pq *ThreadSafePriorityQueue[T]) Dequeue() (T, bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if pq.size == 0 {
		var zero T
		return zero, false
	}

	item := pq.items[0]
	pq.size--
	if pq.size > 0 {
		pq.items[0] = pq.items[pq.size]
		pq.items = pq.items[:pq.size]
		pq.heapifyDown(0)
	} else {
		pq.items = pq.items[:0]
	}

	return item.Value, true
}

// Peek returns the highest priority item without removing it
func (pq *ThreadSafePriorityQueue[T]) Peek() (T, bool) {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	if pq.size == 0 {
		var zero T
		return zero, false
	}

	return pq.items[0].Value, true
}

// Size returns the number of items in the priority queue
func (pq *ThreadSafePriorityQueue[T]) Size() int {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()
	return pq.size
}

// IsEmpty returns true if the priority queue is empty
func (pq *ThreadSafePriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

// Clear removes all items from the priority queue
func (pq *ThreadSafePriorityQueue[T]) Clear() {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	pq.items = pq.items[:0]
	pq.size = 0
}

// heapifyUp maintains the heap property by moving an item up
func (pq *ThreadSafePriorityQueue[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.items[parentIndex].Priority >= pq.items[index].Priority {
			break
		}
		pq.items[parentIndex], pq.items[index] = pq.items[index], pq.items[parentIndex]
		index = parentIndex
	}
}

// heapifyDown maintains the heap property by moving an item down
func (pq *ThreadSafePriorityQueue[T]) heapifyDown(index int) {
	for {
		largest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < pq.size && pq.items[leftChild].Priority > pq.items[largest].Priority {
			largest = leftChild
		}
		if rightChild < pq.size && pq.items[rightChild].Priority > pq.items[largest].Priority {
			largest = rightChild
		}

		if largest == index {
			break
		}

		pq.items[index], pq.items[largest] = pq.items[largest], pq.items[index]
		index = largest
	}
}

// GetPriority returns the priority of an item if it exists in the queue
func (pq *ThreadSafePriorityQueue[T]) GetPriority(value T, equals func(T, T) bool) (int, bool) {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()

	for _, item := range pq.items {
		if equals(item.Value, value) {
			return item.Priority, true
		}
	}
	return 0, false
}

// UpdatePriority updates the priority of an item if it exists in the queue
func (pq *ThreadSafePriorityQueue[T]) UpdatePriority(value T, newPriority int, equals func(T, T) bool) bool {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	for i, item := range pq.items {
		if equals(item.Value, value) {
			oldPriority := item.Priority
			pq.items[i].Priority = newPriority
			if newPriority > oldPriority {
				pq.heapifyUp(i)
			} else {
				pq.heapifyDown(i)
			}
			return true
		}
	}
	return false
}
