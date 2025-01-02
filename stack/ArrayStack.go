package stack

import (
	"fmt"
	"sync"
)

// ArrayStack represents a generic array-based stack
type ArrayStack[T any] struct {
	Arr     []T
	ArrSize int
	Index   int
	mutex   sync.RWMutex
}

// NewArrayStack creates a new generic array-based stack
func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		Arr:     make([]T, 2),
		ArrSize: 2,
		Index:   0,
		mutex:   sync.RWMutex{},
	}
}

// Push adds data to the stack
func (arr *ArrayStack[T]) Push(data T) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.Index >= arr.ArrSize {
		newArr := make([]T, arr.ArrSize*2)
		copy(newArr, arr.Arr)
		arr.Arr = newArr
		arr.ArrSize *= 2
	}
	arr.Arr[arr.Index] = data
	arr.Index++
}

// Pop removes data from the stack
func (arr *ArrayStack[T]) Pop() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.Index == 0 {
		return
	}
	arr.Index--
	var zero T
	arr.Arr[arr.Index] = zero
	if arr.Index <= arr.ArrSize/4 && arr.ArrSize > 2 {
		newArr := make([]T, arr.ArrSize/2)
		copy(newArr, arr.Arr[:arr.Index])
		arr.Arr = newArr
		arr.ArrSize /= 2
	}
}

// IsEmpty returns true if stack is empty
func (arr *ArrayStack[T]) IsEmpty() bool {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	return arr.Index == 0
}

// List returns a slice of stack data
func (arr *ArrayStack[T]) List() []T {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	var list []T
	for i := 0; i < arr.Index; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print displays stack data
func (arr *ArrayStack[T]) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
