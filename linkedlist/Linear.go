package linkedlist

import (
	"fmt"
	"sync"
)

// Linear represents a generic linear linked list
type Linear[T any] struct {
	X     T
	Next  *Linear[T]
	mutex sync.RWMutex
}

// NewLinear creates a new generic linear linked list node
func NewLinear[T any](data T) *Linear[T] {
	return &Linear[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// AddToStart adds data at the beginning of the list
func (node *Linear[T]) AddToStart(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	oldData := node.X
	oldNext := node.Next
	node.X = data
	node.Next = &Linear[T]{X: oldData, Next: oldNext}
}

// AddToSequentially adds data in sorted order
func (node *Linear[T]) AddToSequentially(data T, less func(T, T) bool) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	// Handle the first node
	if node.Next == nil || less(data, node.X) {
		if less(data, node.X) {
			oldData := node.X
			oldNext := node.Next
			node.X = data
			node.Next = &Linear[T]{X: oldData, Next: oldNext, mutex: sync.RWMutex{}}
		} else {
			node.Next = &Linear[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
		}
		return
	}

	// Find the correct position to insert
	iter := node
	for iter.Next != nil && !less(data, iter.Next.X) {
		iter = iter.Next
	}

	// Insert the new node
	newNode := &Linear[T]{X: data, Next: iter.Next, mutex: sync.RWMutex{}}
	iter.Next = newNode
}

// AddToAfter adds data after the specified value
func (node *Linear[T]) AddToAfter(data T, which T, equals func(T, T) bool) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for !equals(iter.X, which) && iter.Next != nil {
		iter = iter.Next
	}
	if equals(iter.X, which) {
		iter.Next = &Linear[T]{X: data, Next: iter.Next, mutex: sync.RWMutex{}}
		return nil
	}
	return fmt.Errorf("value not found")
}

// AddToEnd adds data at the end of the list
func (node *Linear[T]) AddToEnd(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &Linear[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Delete removes data from the list
func (node *Linear[T]) Delete(data T, equals func(T, T) bool) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	if equals(iter.X, data) {
		if iter.Next != nil {
			node.X = iter.Next.X
			node.Next = iter.Next.Next
			return nil
		}
		var zero T
		node.X = zero
		node.Next = nil
		return nil
	}

	for iter.Next != nil && !equals(iter.Next.X, data) {
		iter = iter.Next
	}
	if iter.Next == nil {
		return fmt.Errorf("value not found")
	}
	iter.Next = iter.Next.Next
	return nil
}

// Search looks for data in the list
func (node *Linear[T]) Search(data T, equals func(T, T) bool) bool {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	iter := node
	for iter != nil {
		if equals(iter.X, data) {
			return true
		}
		iter = iter.Next
	}
	return false
}

// List returns a slice of list data
func (node *Linear[T]) List() []T {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	var list []T
	iter := node
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print displays list data
func (node *Linear[T]) Print() {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
