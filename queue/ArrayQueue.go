package queue

import (
	"fmt"
	"sync"
)

// ArrayQueue represents a generic array-based queue
type ArrayQueue[T any] struct {
	Arr        []T
	ArrSize    int
	FirstIndex int
	LastIndex  int
	mutex      sync.RWMutex
}

// NewArrayQueue creates a new generic array-based queue
func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		Arr:        make([]T, 2),
		ArrSize:    2,
		FirstIndex: 0,
		LastIndex:  0,
		mutex:      sync.RWMutex{},
	}
}

// Enqueue adds data to the queue
func (arr *ArrayQueue[T]) Enqueue(data T) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	// Removing an element from the array is just row shifting if before enlarging the array.
	// If the first index is different from 0, the element has been removed from the array
	// Then we will reorder the array instead of enlarging the array.
	// If first index is 0 and last index is bigger than array size we will increase array size
	if arr.LastIndex >= arr.ArrSize {
		if arr.FirstIndex == 0 {
			newArr := make([]T, arr.ArrSize*2)
			copy(newArr, arr.Arr)
			arr.Arr = newArr
			arr.ArrSize *= 2
		} else {
			arr.reSort()
		}
	}
	arr.Arr[arr.LastIndex] = data
	arr.LastIndex++
}

// Dequeue removes data from the queue
func (arr *ArrayQueue[T]) Dequeue() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	// if deque is run first
	if arr.FirstIndex == 0 && arr.LastIndex == 0 {
		return
	}
	var zero T
	arr.Arr[arr.FirstIndex] = zero
	arr.FirstIndex++
	if arr.LastIndex-arr.FirstIndex <= arr.ArrSize/4 {
		newArr := make([]T, arr.ArrSize/2)
		sort := 0
		for i := arr.FirstIndex; i < arr.LastIndex; i++ {
			newArr[sort] = arr.Arr[i]
			sort++
		}

		arr.Arr = newArr
		arr.ArrSize /= 2
		arr.LastIndex = arr.LastIndex - arr.FirstIndex
		arr.FirstIndex = 0
	}
}

// reSort reorders the queue data
func (arr *ArrayQueue[T]) reSort() {
	newArr := make([]T, arr.ArrSize)
	sort := 0
	for i := arr.FirstIndex; i < arr.LastIndex; i++ {
		newArr[sort] = arr.Arr[i]
		sort++
	}
	arr.Arr = newArr
	arr.LastIndex = arr.LastIndex - arr.FirstIndex
	arr.FirstIndex = 0
}

// List returns a slice of queue data
func (arr *ArrayQueue[T]) List() []T {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()

	if arr.FirstIndex >= arr.LastIndex {
		return []T{}
	}
	var list []T
	for i := arr.FirstIndex; i < arr.LastIndex; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print displays queue data
func (arr *ArrayQueue[T]) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
