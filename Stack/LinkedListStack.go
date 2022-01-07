package Stack

import "fmt"

type ILinkedListStack interface{
	Push(data int)
	Pop()
	Print()
	List() []int
}

type linkedListStack struct {
	X int
	Next *linkedListStack
}

func LinkedListStack(data int) ILinkedListStack{
	return &linkedListStack{data, nil}
}

// Push Add to data
func (arr *linkedListStack) Push(data int) {
	iter := arr
	if iter.X == -1{
		iter.X = data
	}else{
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = &linkedListStack{X: data, Next: nil}
	}
}

// Pop Remove to data
func (arr *linkedListStack) Pop() {
	iter := arr
	if iter.Next != nil{
		for iter.Next.Next != nil {
			iter = iter.Next
		}
		iter.Next = nil
	}else{
		arr.X = -1
	}
}

// List - data slice
func (arr *linkedListStack) List() []int{
	var list []int
	iter := arr
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print data
func (arr *linkedListStack) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val," ")
	}
	fmt.Println()
}