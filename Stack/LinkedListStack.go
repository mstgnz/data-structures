package stack

import "fmt"

type ILinkedListStack interface {
	Push(data int)
	Pop()
	IsEmpty() bool
	Print()
	List() []int
}

type linkedListStack struct {
	X    int
	Next *linkedListStack
}

func LinkedListStack(data int) ILinkedListStack {
	return &linkedListStack{data, nil}
}

// Push Add to data at the beginning (LIFO)
func (arr *linkedListStack) Push(data int) {
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

// Pop Remove to data from the beginning
func (arr *linkedListStack) Pop() {
	if arr.IsEmpty() {
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
	return arr.X == -1 && arr.Next == nil
}

// List - data slice
func (arr *linkedListStack) List() []int {
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

// Print data
func (arr *linkedListStack) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
