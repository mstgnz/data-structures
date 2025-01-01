package linkedlist

import (
	"fmt"
	"sync"
)

type ILinear interface {
	AddToStart(data int)
	AddToSequentially(data int)
	AddToAfter(data, which int) error
	AddToEnd(data int)
	Delete(data int) error
	Search(data int) bool
	List() []int
	Print()
}

type linear struct {
	X     int
	Next  *linear
	mutex sync.RWMutex
}

func Linear(data int) ILinear {
	return &linear{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// AddToStart adds data at the beginning of the list
func (node *linear) AddToStart(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	oldData := node.X
	oldNext := node.Next
	node.X = data
	node.Next = &linear{X: oldData, Next: oldNext}
}

// AddToSequentially adds data in sorted order
func (node *linear) AddToSequentially(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	if node.X > data {
		// If the new data is smaller than the current node's data,
		// insert it at the beginning
		oldData := node.X
		oldNext := node.Next
		node.X = data
		node.Next = &linear{X: oldData, Next: oldNext}
		return
	}

	iter := node
	for iter.Next != nil && iter.Next.X < data {
		iter = iter.Next
	}
	iter.Next = &linear{X: data, Next: iter.Next, mutex: sync.RWMutex{}}
}

// AddToAfter adds data after the specified value
func (node *linear) AddToAfter(data int, which int) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.X != which && iter.Next != nil {
		iter = iter.Next
	}
	if iter.X == which {
		iter.Next = &linear{X: data, Next: iter.Next, mutex: sync.RWMutex{}}
		return nil
	}
	return fmt.Errorf("%d not found", which)
}

// AddToEnd adds data at the end of the list
func (node *linear) AddToEnd(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &linear{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Delete removes data from the list
func (node *linear) Delete(data int) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	if iter.X == data {
		if iter.Next != nil {
			node.X = iter.Next.X
			node.Next = iter.Next.Next
			return nil
		}
		node.X = 0
		node.Next = nil
		return nil
	}

	for iter.Next != nil && iter.Next.X != data {
		iter = iter.Next
	}
	if iter.Next == nil {
		return fmt.Errorf("%d not found", data)
	}
	iter.Next = iter.Next.Next
	return nil
}

// Search looks for data in the list
func (node *linear) Search(data int) bool {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	iter := node
	for iter != nil {
		if iter.X == data {
			return true
		}
		iter = iter.Next
	}
	return false
}

// List returns a slice of list data
func (node *linear) List() []int {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	var list []int
	iter := node
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print displays list data
func (node *linear) Print() {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
