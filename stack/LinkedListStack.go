package stack

import (
	"fmt"
	"sync"
)

// LinkedListStack represents a generic linked list-based stack
type LinkedListStack[T comparable] struct {
	X     T
	Next  *LinkedListStack[T]
	mutex sync.RWMutex
}

// NewLinkedListStack creates a new generic linked list-based stack
func NewLinkedListStack[T comparable](data T) *LinkedListStack[T] {
	return &LinkedListStack[T]{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Push adds data at the beginning (LIFO)
func (arr *LinkedListStack[T]) Push(data T) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	newNode := &LinkedListStack[T]{X: data, Next: nil}
	newNode.Next = arr.Next
	arr.Next = newNode
	temp := arr.X
	arr.X = data
	newNode.X = temp
}

// Pop removes data from the beginning
func (arr *LinkedListStack[T]) Pop() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	var zero T
	if arr.X == zero && arr.Next == nil {
		return
	}
	if arr.Next == nil {
		arr.X = zero
		return
	}
	arr.X = arr.Next.X
	arr.Next = arr.Next.Next
}

// IsEmpty returns true if stack is empty
func (arr *LinkedListStack[T]) IsEmpty() bool {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	var zero T
	return arr.X == zero && arr.Next == nil
}

// List returns a slice of stack data
func (arr *LinkedListStack[T]) List() []T {
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

// Print displays stack data
func (arr *LinkedListStack[T]) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
