package linkedlist

import (
	"fmt"
	"sync"
)

// Circular represents a generic circular linked list
type Circular[T any] struct {
	X     T
	Next  *Circular[T]
	mutex sync.RWMutex
}

// NewCircular creates a new generic circular linked list node
func NewCircular[T any](data T) *Circular[T] {
	init := &Circular[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
	init.Next = init
	return init
}

// AddToStart adds data at the beginning of the list
func (node *Circular[T]) AddToStart(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	oldData := node.X
	oldNext := node.Next
	node.X = data
	node.Next = &Circular[T]{X: oldData, Next: oldNext}
}

// AddToSequentially adds data in sorted order
func (node *Circular[T]) AddToSequentially(data T, less func(T, T) bool) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	// If the value to be added is less than the value of the root object
	if less(data, node.X) {
		temp := node.X
		node.X = data
		newNode := &Circular[T]{X: temp, Next: node.Next}
		node.Next = newNode
		for iter.Next != node {
			iter = iter.Next
		}
		iter.Next = node
	} else {
		// Advance up to the value that is less than the value you want to add.
		for iter.Next != node && less(iter.Next.X, data) {
			iter = iter.Next
		}
		// Add the value to the next of the object that is smaller than the value to be added, by creating a new object.
		// add the current next to the next of the newly added object
		iter.Next = &Circular[T]{X: data, Next: iter.Next}
	}
}

// AddToAfter adds data after the specified value
func (node *Circular[T]) AddToAfter(data T, which T, equals func(T, T) bool) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node

	// Check all nodes
	for {
		if equals(iter.X, which) {
			newNode := &Circular[T]{X: data, Next: iter.Next}
			iter.Next = newNode
			return
		}
		iter = iter.Next
		// Returned to start and not found
		if iter == node {
			fmt.Println("value not found!")
			return
		}
	}
}

// AddToEnd adds data at the end of the list
func (node *Circular[T]) AddToEnd(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != node {
		iter = iter.Next
	}
	iter.Next = &Circular[T]{X: data, Next: node}
}

// Delete removes data from the list
func (node *Circular[T]) Delete(data T, equals func(T, T) bool) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	// If the value to be deleted is a value in between or at the end, we move our iter object to the previous node object to be deleted.
	iter := node
	// If the root object is to be deleted
	if equals(iter.X, data) {
		if node.Next == node {
			// If it's the only element in the list
			var zero T
			node.X = zero
		} else {
			for iter.Next != node {
				iter = iter.Next
			}
			node.X = node.Next.X
			node.Next = node.Next.Next
			iter.Next = node
		}
		return nil
	}

	// If one of the other elements is wanted to be deleted
	for iter.Next != node && !equals(iter.Next.X, data) {
		iter = iter.Next
	}
	if iter.Next == node {
		return fmt.Errorf("value not found")
	}
	iter.Next = iter.Next.Next
	return nil
}

// List returns a slice of list data
func (node *Circular[T]) List() []T {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	var list []T
	iter := node
	list = append(list, iter.X)
	iter = iter.Next
	for iter != node {
		list = append(list, iter.X)

		iter = iter.Next
	}
	return list
}

// Print displays list data
func (node *Circular[T]) Print() {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
