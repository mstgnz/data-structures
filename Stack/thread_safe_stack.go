package stack

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// ThreadSafeStack represents a thread-safe generic stack data structure
type ThreadSafeStack[T utils.Any] struct {
	items []T
	mutex sync.RWMutex
}

// NewThreadSafeStack creates a new thread-safe stack
func NewThreadSafeStack[T utils.Any]() *ThreadSafeStack[T] {
	return &ThreadSafeStack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *ThreadSafeStack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *ThreadSafeStack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// Peek returns the top item without removing it
func (s *ThreadSafeStack[T]) Peek() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *ThreadSafeStack[T]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *ThreadSafeStack[T]) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.items)
}

// Clear removes all items from the stack
func (s *ThreadSafeStack[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = make([]T, 0)
}
