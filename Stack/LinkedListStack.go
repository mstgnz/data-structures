package stack

import (
	"fmt"
	"sync"
)

type ILinkedListStack interface {
	Push(data int)
	Pop()
	IsEmpty() bool
	Print()
	List() []int
}

type linkedListStack struct {
	X     int
	Next  *linkedListStack
	mutex sync.RWMutex
}

func LinkedListStack(data int) ILinkedListStack {
	return &linkedListStack{X: data, Next: nil, mutex: sync.RWMutex{}}
}

// Push adds data at the beginning (LIFO)
func (arr *linkedListStack) Push(data int) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.X == -1 {
		arr.X = data
		return
	}
	newNode := &linkedListStack{X: data, Next: nil}
	newNode.Next = arr.Next
	arr.Next = newNode
	temp := arr.X
	arr.X = data
	newNode.X = temp
}

// Pop removes data from the beginning
func (arr *linkedListStack) Pop() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.X == -1 && arr.Next == nil {
		return
	}
	if arr.Next == nil {
		arr.X = -1
		return
	}
	arr.X = arr.Next.X
	arr.Next = arr.Next.Next
}

// IsEmpty returns true if stack is empty
func (arr *linkedListStack) IsEmpty() bool {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	return arr.X == -1 && arr.Next == nil
}

// List returns a slice of stack data
func (arr *linkedListStack) List() []int {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	var list []int
	iter := arr
	for iter != nil {
		if iter.X != -1 {
			list = append(list, iter.X)
		}
		iter = iter.Next
	}
	return list
}

// Print displays stack data
func (arr *linkedListStack) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
