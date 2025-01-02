package queue

import (
	"fmt"
	"sync"
)

// LinkedListQueue represents a generic linked list-based queue
type LinkedListQueue[T comparable] struct {
	X     T
	Next  *LinkedListQueue[T]
	mutex sync.RWMutex
}

// NewLinkedListQueue creates a new generic linked list-based queue
func NewLinkedListQueue[T comparable](data T) *LinkedListQueue[T] {
	return &LinkedListQueue[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Enqueue adds data to the queue
func (arr *LinkedListQueue[T]) Enqueue(data T) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	iter := arr
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &LinkedListQueue[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Dequeue removes data from the queue
func (arr *LinkedListQueue[T]) Dequeue() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	var zero T
	if arr.X == zero && arr.Next == nil {
		return
	}
	if arr.Next != nil {
		arr.X = arr.Next.X
		arr.Next = arr.Next.Next
	} else {
		arr.X = zero
	}
}

// List returns a slice of queue data
func (arr *LinkedListQueue[T]) List() []T {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()

	var list []T
	iter := arr
	list = append(list, iter.X)
	for iter.Next != nil {
		iter = iter.Next
		list = append(list, iter.X)
	}
	return list
}

// Print displays queue data
func (arr *LinkedListQueue[T]) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
