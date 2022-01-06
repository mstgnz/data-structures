package Queue

import "fmt"

type ILinkedListQueue interface{
	Enqueue(data int)
	Dequeue()
	List() []int
	Print()
}

type linkedListQueue struct {
	X int
	Next *linkedListQueue
}

func LinkedListQueue(data int) ILinkedListQueue{
	return &linkedListQueue{data, nil}
}

// Enqueue Add to data
func (arr *linkedListQueue) Enqueue(data int) {
	iter := arr
	if iter.X == -1{
		iter.X = data
	}else{
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = &linkedListQueue{X: data, Next: nil}
	}
}

// Dequeue Remove to data
func (arr *linkedListQueue) Dequeue() {
	if arr.Next != nil{
		*arr = *arr.Next
	}else{
		arr.X = -1
	}
}

// List data - slice
func (arr *linkedListQueue) List() []int{
	var list []int
	iter := arr
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print data
func (arr *linkedListQueue) Print() {
	fmt.Print("print : ")
	for _, val := range arr.List() {
		fmt.Print(val," ")
	}
	fmt.Println()
}