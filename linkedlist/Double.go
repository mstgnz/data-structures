package linkedlist

import (
	"fmt"
	"sync"
)

// Double represents a generic double linked list
type Double[T any] struct {
	X     T
	Next  *Double[T]
	Prev  *Double[T]
	mutex sync.RWMutex
}

// NewDouble creates a new generic double linked list node
func NewDouble[T any](data T) *Double[T] {
	return &Double[T]{X: data, Next: nil, Prev: nil, mutex: sync.RWMutex{}}
}

// AddToStart adds data at the beginning of the list
func (node *Double[T]) AddToStart(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	oldData := node.X
	oldNext := node.Next
	node.X = data
	node.Next = &Double[T]{X: oldData, Next: oldNext, Prev: node}
	if node.Next.Next != nil {
		node.Next.Next.Prev = node.Next
	}
}

// AddToSequentially adds data in sorted order
func (node *Double[T]) AddToSequentially(data T, less func(T, T) bool) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil && less(iter.Next.X, data) {
		iter = iter.Next
	}
	newNode := &Double[T]{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
	iter.Next = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// AddToAfter adds data after the specified value
func (node *Double[T]) AddToAfter(data T, which T, equals func(T, T) bool) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	found := false

	// Check the first node
	if equals(iter.X, which) {
		found = true
		newNode := &Double[T]{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
		if iter.Next != nil {
			iter.Next.Prev = newNode
		}
		iter.Next = newNode
		return
	}

	// Check other nodes
	for iter.Next != nil {
		iter = iter.Next
		if equals(iter.X, which) {
			found = true
			newNode := &Double[T]{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
			if iter.Next != nil {
				iter.Next.Prev = newNode
			}
			iter.Next = newNode
			return
		}
	}

	if !found {
		fmt.Println("value not found!")
	}
}

// AddToEnd adds data at the end of the list
func (node *Double[T]) AddToEnd(data T) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &Double[T]{X: data, Next: nil, Prev: iter, mutex: sync.RWMutex{}}
}

// Delete removes data from the list
func (node *Double[T]) Delete(data T, equals func(T, T) bool) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	// If the value to be deleted is the first element
	if equals(node.X, data) {
		if node.Next != nil {
			node.X = node.Next.X
			node.Next = node.Next.Next
			if node.Next != nil {
				node.Next.Prev = node
			}
		} else {
			var zero T
			node.X = zero
		}
		return nil
	}

	// If the value to be deleted is a value in between or at the end
	iter := node
	for iter.Next != nil && !equals(iter.Next.X, data) {
		iter = iter.Next
	}
	if iter.Next == nil {
		return fmt.Errorf("value not found")
	}

	// Delete the node
	iter.Next = iter.Next.Next
	if iter.Next != nil {
		iter.Next.Prev = iter
	}
	return nil
}

// List returns a slice of list data
func (node *Double[T]) List(reverse bool) []T {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	var list []T
	iter := node
	if reverse { // print bottom to top
		for iter.Next != nil {
			iter = iter.Next
		}
		for iter != nil {
			list = append(list, iter.X)
			iter = iter.Prev
		}
	} else { // print top to bottom
		for iter != nil {
			list = append(list, iter.X)
			iter = iter.Next
		}
	}
	return list
}

// Print displays list data
func (node *Double[T]) Print(reverse bool) {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range node.List(reverse) {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
