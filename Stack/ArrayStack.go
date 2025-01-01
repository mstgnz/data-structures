package stack

import "fmt"

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
}

func ArrayStack() IArrayStack {
	return &arrayStack{[]int{0, 0}, 2, 0}
}

// Push Add to data
func (arr *arrayStack) Push(data int) {
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

// Pop Remove to data
func (arr *arrayStack) Pop() {
	if arr.IsEmpty() {
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
	return arr.Index == 0
}

// List - data slice
func (arr *arrayStack) List() []int {
	var list []int
	for i := 0; i < arr.Index; i++ {
		list = append(list, arr.Arr[i])
	}
	return list
}

// Print data
func (arr *arrayStack) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
