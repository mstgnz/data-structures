package LinkedList

import "fmt"

type ICircular interface{
	AddToStart(data int)
	AddToSequentially(data int)
	AddToAfter(data, which int)
	AddToEnd(data int)
	Delete(data int)
	List() []int
	Print()
}

type circular struct {
	X    int
	Next *circular
}

func Circular(data int) ICircular{
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
	for node.X != which && node != node.Next {
		node = node.Next
	}
	if node.X == which{
		temp := *node
		node.Next = &circular{X: data, Next: nil}
		node.Next.Next = temp.Next
	}else{
		fmt.Println(which,"not found!")
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
func (node *circular) Delete(data int) {
	// If the value to be deleted is a value in between or at the end, we move our iter object to the previous node object to be deleted.
	iter := node
	// If the root object is to be deleted
	if iter.X == data {
		for iter.Next != node {
			iter = iter.Next
		}
		node.X = node.Next.X
		node.Next = node.Next.Next
		iter.Next = node
		// If one of the other elements is wanted to be deleted
	} else {
		for iter.Next != node && iter.Next.X != data {
			iter = iter.Next
		}
		if iter.Next == node {
			fmt.Println(data,"not found")
		} else {
			iter.Next = iter.Next.Next
		}
	}
}

// List data - slice
func (node *circular) List() []int{
	var list []int
	iter := node
	list = append(list, iter.X)
	iter = iter.Next
	for iter != node{
		list = append(list, iter.X)
		iter = iter.Next
	}
	return list
}

// Print data
func (node *circular) Print() {
	fmt.Print("print : ")
	for _, val := range node.List() {
		fmt.Print(val," ")
	}
	fmt.Println()
}