package linkedlist

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// Node represents a node in the linked list
type Node[T utils.Any] struct {
	Value T
	Next  *Node[T]
}

// ThreadSafeLinkedList represents a thread-safe generic linked list
type ThreadSafeLinkedList[T utils.Any] struct {
	head  *Node[T]
	size  int
	mutex sync.RWMutex
}

// NewThreadSafeLinkedList creates a new thread-safe linked list
func NewThreadSafeLinkedList[T utils.Any]() *ThreadSafeLinkedList[T] {
	return &ThreadSafeLinkedList[T]{
		head: nil,
		size: 0,
	}
}

// AddFirst adds an item to the beginning of the list
func (l *ThreadSafeLinkedList[T]) AddFirst(value T) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newNode := &Node[T]{Value: value, Next: l.head}
	l.head = newNode
	l.size++
}

// AddLast adds an item to the end of the list
func (l *ThreadSafeLinkedList[T]) AddLast(value T) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

// RemoveFirst removes and returns the first item
func (l *ThreadSafeLinkedList[T]) RemoveFirst() (T, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var zero T
	if l.head == nil {
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.Next
	l.size--
	return value, true
}

// RemoveLast removes and returns the last item
func (l *ThreadSafeLinkedList[T]) RemoveLast() (T, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var zero T
	if l.head == nil {
		return zero, false
	}

	if l.head.Next == nil {
		value := l.head.Value
		l.head = nil
		l.size--
		return value, true
	}

	current := l.head
	for current.Next.Next != nil {
		current = current.Next
	}
	value := current.Next.Value
	current.Next = nil
	l.size--
	return value, true
}

// GetFirst returns the first item without removing it
func (l *ThreadSafeLinkedList[T]) GetFirst() (T, bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var zero T
	if l.head == nil {
		return zero, false
	}
	return l.head.Value, true
}

// GetLast returns the last item without removing it
func (l *ThreadSafeLinkedList[T]) GetLast() (T, bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var zero T
	if l.head == nil {
		return zero, false
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	return current.Value, true
}

// Contains checks if an item exists in the list
// This method only works with comparable types
func Contains[T comparable](l *ThreadSafeLinkedList[T], value T) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	current := l.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// IsEmpty returns true if the list is empty
func (l *ThreadSafeLinkedList[T]) IsEmpty() bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.head == nil
}

// Size returns the number of items in the list
func (l *ThreadSafeLinkedList[T]) Size() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.size
}

// Clear removes all items from the list
func (l *ThreadSafeLinkedList[T]) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.head = nil
	l.size = 0
}
