package LinkedList

import "fmt"

type ILinear interface{
	AddToStart(data int)
	AddToSequentially(data int)
	AddToBetween(data, which int)
	AddToEnd(data int)
	Delete(data int)
	List() []int
	Print()
}

type linear struct {
	X    int
	Next *linear
}

func Linear(data int) ILinear{
	return &linear{X: data}
}

// AddToStart data
func (node *linear) AddToStart(data int) {
	temp := *node
	node.X = data
	node.Next = &temp
}

// AddToSequentially data
func (node *linear) AddToSequentially(data int) {
	for node.Next != nil && node.Next.X < data {
		node = node.Next
	}
	temp := *node
	node.Next = &linear{X: data, Next: nil}
	node.Next.Next = temp.Next
}

// AddToBetween data
func (node *linear) AddToBetween(data int, which int) {
	for node.X != which && node.Next != nil {
		node = node.Next
	}
	if node.X == which{
		temp := *node
		node.Next = &linear{X: data, Next: nil}
		node.Next.Next = temp.Next
	}else{
		fmt.Println(which,"not found!")
	}
}

// AddToEnd data
func (node *linear) AddToEnd(data int) {
	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &linear{X: data, Next: nil}
}

// Delete data
func (node *linear) Delete(data int) {
	iter := node
	if iter.X == data{
		if iter.Next != nil {
			node.X = iter.Next.X
			node.Next = iter.Next.Next
		}else{
			fmt.Println(data,"is set to zero because it is the last element.")
			node.X = 0
		}
	}else{
		for iter.Next != nil && iter.Next.X != data {
			iter = iter.Next
		}
		if iter.Next == nil {
			fmt.Println(data,"not found!")
		} else {
			node.Next = iter.Next.Next
		}
	}
}

// List data - slice
func (node *linear) List() []int{
	var list []int
	iter := node
	for iter != nil {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print data
func (node *linear) Print() {
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val," ")
	}
	fmt.Println()
}
