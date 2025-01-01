package stack

import (
	"fmt"
	"sync"
)

type IArrayStack interface {
	Push(data int)
	Pop()
	IsEmpty() bool
	Print()
	List() []int
}

type arrayStack struct {
	Arr     []int
	ArrSize int
	Index   int
	mutex   sync.RWMutex
}

func ArrayStack() IArrayStack {
	return &arrayStack{
		Arr:     []int{0, 0},
		ArrSize: 2,
		Index:   0,
		mutex:   sync.RWMutex{},
	}
}

// Push adds data to the stack
func (arr *arrayStack) Push(data int) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.Index >= arr.ArrSize {
		newArr := make([]int, arr.ArrSize*2)
		for i := 0; i < arr.ArrSize; i++ {
			newArr[i] = arr.Arr[i]
		}
		arr.Arr = newArr
		arr.ArrSize *= 2
	}
	arr.Arr[arr.Index] = data
	arr.Index++
}

// Pop removes data from the stack
func (arr *arrayStack) Pop() {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	if arr.Index == 0 {
		return
	}
	arr.Index--
	arr.Arr[arr.Index] = 0
	if arr.Index <= arr.ArrSize/4 && arr.ArrSize > 2 {
		newArr := make([]int, arr.ArrSize/2)
		for i := 0; i < arr.Index; i++ {
			newArr[i] = arr.Arr[i]
		}
		arr.Arr = newArr
		arr.ArrSize /= 2
	}
}

// IsEmpty returns true if stack is empty
func (arr *arrayStack) IsEmpty() bool {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	return arr.Index == 0
}

// List returns a slice of stack data
func (arr *arrayStack) List() []int {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	var list []int
	for i := 0; i < arr.Index; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print displays stack data
func (arr *arrayStack) Print() {
	arr.mutex.RLock()
	defer arr.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
