package linkedlist

import (
	"fmt"
	"sync"
)

type Double struct {
	X     int
	Next  *Double
	Prev  *Double
	mutex sync.RWMutex
}

func NewDouble(data int) *Double {
	return &Double{X: data, Next: nil, Prev: nil, mutex: sync.RWMutex{}}
}

// AddToStart adds data at the beginning of the list
func (node *Double) AddToStart(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	oldData := node.X
	oldNext := node.Next
	node.X = data
	node.Next = &Double{X: oldData, Next: oldNext, Prev: node}
	if node.Next.Next != nil {
		node.Next.Next.Prev = node.Next
	}
}

// AddToSequentially adds data in sorted order
func (node *Double) AddToSequentially(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil && iter.Next.X < data {
		iter = iter.Next
	}
	newNode := &Double{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
	iter.Next = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// AddToAfter adds data after the specified value
func (node *Double) AddToAfter(data int, which int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	found := false

	// Check the first node
	if iter.X == which {
		found = true
		newNode := &Double{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
		if iter.Next != nil {
			iter.Next.Prev = newNode
		}
		iter.Next = newNode
		return
	}

	// Check other nodes
	for iter.Next != nil {
		iter = iter.Next
		if iter.X == which {
			found = true
			newNode := &Double{X: data, Next: iter.Next, Prev: iter, mutex: sync.RWMutex{}}
			if iter.Next != nil {
				iter.Next.Prev = newNode
			}
			iter.Next = newNode
			return
		}
	}

	if !found {
		fmt.Println(which, "not found!")
	}
}

// AddToEnd adds data at the end of the list
func (node *Double) AddToEnd(data int) {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &Double{X: data, Next: nil, Prev: iter, mutex: sync.RWMutex{}}
}

// Delete removes data from the list
func (node *Double) Delete(data int) error {
	node.mutex.Lock()
	defer node.mutex.Unlock()

	// If the value to be deleted is the first element
	if node.X == data {
		if node.Next != nil {
			node.X = node.Next.X
			node.Next = node.Next.Next
			if node.Next != nil {
				node.Next.Prev = node
			}
		} else {
			node.X = 0
		}
		return nil
	}

	// If the value to be deleted is a value in between or at the end
	iter := node
	for iter.Next != nil && iter.Next.X != data {
		iter = iter.Next
	}
	if iter.Next == nil {
		return fmt.Errorf("%d not found", data)
	}

	// Delete the node
	iter.Next = iter.Next.Next
	if iter.Next != nil {
		iter.Next.Prev = iter
	}
	return nil
}

// List returns a slice of list data
func (node *Double) List(reverse bool) []int {
	node.mutex.RLock()
	defer node.mutex.RUnlock()

	var list []int
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
func (node *Double) Print(reverse bool) {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range node.List(reverse) {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
