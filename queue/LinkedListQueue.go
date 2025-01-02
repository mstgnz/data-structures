package queue

import (
	"fmt"
	"sync"
)

type LinkedListQueue struct {
	X     int
	Next  *LinkedListQueue
	mutex sync.RWMutex
}

func NewLinkedListQueue(data int) *LinkedListQueue {
	return &LinkedListQueue{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Enqueue adds data to the queue
func (arr *LinkedListQueue) Enqueue(data int) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	iter := arr
	if iter.X == -1 {
		iter.X = data
	} else {
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = &LinkedListQueue{X: data, Next: nil, mutex: sync.RWMutex{}}
	}
}

// Dequeue removes data from the queue
func (arr *LinkedListQueue) Dequeue() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.X == -1 && arr.Next == nil {
		return
	}
	if arr.Next != nil {
		arr.X = arr.Next.X
		arr.Next = arr.Next.Next
	} else {
		arr.X = -1
	}
}

// List returns a slice of queue data
func (arr *LinkedListQueue) List() []int {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()

	var list []int
	iter := arr
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print displays queue data
func (arr *LinkedListQueue) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
