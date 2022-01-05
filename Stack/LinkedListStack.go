package Stack

import "fmt"

type ILinkedListStack interface{
	Push(data int)
	Pop()
	Print()
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

// Print data
func (arr *linkedListStack) Print() {
	iter := arr
	for iter != nil {
		fmt.Printf("%v ", iter.X)
		iter = iter.Next
	}
	fmt.Println()
}