package linkedlist

import "fmt"

type ILinear interface {
	AddToStart(data int)
	AddToSequentially(data int)
	AddToAfter(data, which int) error
	AddToEnd(data int)
	Delete(data int) error
	Search(data int) bool
	List() []int
	Print()
}

type linear struct {
	X    int
	Next *linear
}

func Linear(data int) ILinear {
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
	if node.X > data {
		// If the new data is smaller than the current node's data,
		// insert it at the beginning
		temp := *node
		node.X = data
		node.Next = &temp
		return
	}

	iter := node
	for iter.Next != nil && iter.Next.X < data {
		iter = iter.Next
	}
	iter.Next = &linear{X: data, Next: iter.Next}
}

// AddToAfter data
func (node *linear) AddToAfter(data int, which int) error {
	for node.X != which && node.Next != nil {
		node = node.Next
	}
	if node.X == which {
		temp := *node
		node.Next = &linear{X: data, Next: nil}
		node.Next.Next = temp.Next
		return nil
	}
	return fmt.Errorf("%d not found", which)
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
func (node *linear) Delete(data int) error {
	iter := node
	if iter.X == data {
		if iter.Next != nil {
			node.X = iter.Next.X
			node.Next = iter.Next.Next
			return nil
		}
		node.X = 0
		node.Next = nil
		return nil
	}

	for iter.Next != nil && iter.Next.X != data {
		iter = iter.Next
	}
	if iter.Next == nil {
		return fmt.Errorf("%d not found", data)
	}
	iter.Next = iter.Next.Next
	return nil
}

// Search data
func (node *linear) Search(data int) bool {
	iter := node
	for iter != nil {
		if iter.X == data {
			return true
		}
		iter = iter.Next
	}
	return false
}

// List data - slice
func (node *linear) List() []int {
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
		fmt.Print(val, " ")
	}
	fmt.Println()
}
