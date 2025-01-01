package linkedlist

import (
	"fmt"
)

type ICircular interface {
	AddToStart(data int)
	AddToSequentially(data int)
	AddToAfter(data, which int)
	AddToEnd(data int)
	Delete(data int) error
	List() []int
	Print()
}

type circular struct {
	X    int
	Next *circular
}

func Circular(data int) ICircular {
	init := &circular{X: data, Next: nil}
	init.Next = init
	return init
}

// AddToStart data
func (node *circular) AddToStart(data int) {
	temp := *node
	node.X = data
	node.Next = &temp
}

// AddToSequentially data
func (node *circular) AddToSequentially(data int) {
	iter := node
	// If the value to be added is less than the value of the root object
	if node.X > data {
		node.AddToStart(data)
		for iter.Next != node {
			iter = iter.Next
		}
		iter.Next = node
		// If the value to be added is greater than the value of the root object
	} else {
		// Advance up to the value that is less than the value you want to add.
		for iter.Next != node && iter.Next.X < data {
			iter = iter.Next
		}
		// Add the value to the next of the object that is smaller than the value to be added, by creating a new object.
		// add the current next to the next of the newly added object
		iter.Next = &circular{X: data, Next: iter.Next}
	}
}

// AddToAfter data
func (node *circular) AddToAfter(data int, which int) {
	iter := node

	// Check all nodes
	for {
		if iter.X == which {
			newNode := &circular{X: data, Next: iter.Next}
			iter.Next = newNode
			return
		}
		iter = iter.Next
		// Returned to start and not found
		if iter == node {
			fmt.Println(which, "not found!")
			return
		}
	}
}

// AddToEnd data
func (node *circular) AddToEnd(data int) {
	iter := node
	for iter.Next != node {
		iter = iter.Next
	}
	iter.Next = &circular{X: data, Next: node}
}

// Delete data
func (node *circular) Delete(data int) error {
	// If the value to be deleted is a value in between or at the end, we move our iter object to the previous node object to be deleted.
	iter := node
	// If the root object is to be deleted
	if iter.X == data {
		if node.Next == node {
			// If it's the only element in the list
			node.X = 0
		} else {
			for iter.Next != node {
				iter = iter.Next
			}
			node.X = node.Next.X
			node.Next = node.Next.Next
			iter.Next = node
		}
		return nil
	}

	// If one of the other elements is wanted to be deleted
	for iter.Next != node && iter.Next.X != data {
		iter = iter.Next
	}
	if iter.Next == node {
		return fmt.Errorf("%d not found", data)
	}
	iter.Next = iter.Next.Next
	return nil
}

// List data - slice
func (node *circular) List() []int {
	var list []int
	iter := node
	list = append(list, iter.X)
	iter = iter.Next
	for iter != node {
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print data
func (node *circular) Print() {
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
