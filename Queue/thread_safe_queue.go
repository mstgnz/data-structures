package queue

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// ThreadSafeQueue represents a thread-safe generic queue data structure
type ThreadSafeQueue[T utils.Any] struct {
	items []T
	mutex sync.RWMutex
}

// NewThreadSafeQueue creates a new thread-safe queue
func NewThreadSafeQueue[T utils.Any]() *ThreadSafeQueue[T] {
	return &ThreadSafeQueue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the end of the queue
func (q *ThreadSafeQueue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item from the queue
func (q *ThreadSafeQueue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the first item without removing it
func (q *ThreadSafeQueue[T]) Front() (T, bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	return q.items[0], true
}

// IsEmpty returns true if the queue is empty
func (q *ThreadSafeQueue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *ThreadSafeQueue[T]) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.items)
}

// Clear removes all items from the queue
func (q *ThreadSafeQueue[T]) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = make([]T, 0)
}
