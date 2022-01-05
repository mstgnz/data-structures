package LinkedList

import "fmt"

type IDouble interface{
	AddToStart(data int)
	AddToSequentially(data int)
	AddToBetween(data, which int)
	AddToEnd(data int)
	Delete(data int)
	Print(reverse bool)
}

type double struct {
	X    int
	Next *double
	Prev *double
}

func Double(data int) IDouble{
	return &double{X: data}
}

// AddToStart data
func (node *double) AddToStart(data int) {
	temp := *node
	node.X = data
	node.Next = &temp
	node.Next.Prev = node
	if node.Next.Next != nil {
		node.Next.Next.Prev = node.Next
	}
}

// AddToSequentially data
func (node *double) AddToSequentially(data int) {
	for node.Next != nil && node.Next.X < data {
		node = node.Next
	}
	node.Next = &double{X: data, Next: node.Next, Prev: node}
	if node.Next.Next != nil {
		node.Next.Next.Prev = node.Next
	}
}

// AddToBetween data
func (node *double) AddToBetween(data int, which int) {
	for node.Next != nil && node.X != which {
		node = node.Next
	}
	node.Next = &double{X: data, Next: node.Next, Prev: node}
	if node.Next.Next != nil {
		node.Next.Next.Prev = node.Next
	}
}

// AddToEnd data
func (node *double) AddToEnd(data int) {
	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &double{X: data, Next: nil, Prev: iter}
}

// Delete data
func (node *double) Delete(data int) {
	// If the value to be deleted is a value in between or at the end, we move our iter object to the previous node object to be deleted.
	for node.Next != nil && node.Next.X != data {
		node = node.Next
	}
	if node.Next == nil {
		fmt.Println(data,"not found!")
	} else {
		node.Next = node.Next.Next
		node.Next.Prev = node
	}
}

// Print data
func (node *double) Print(reverse bool) {
	iter := node
	if reverse { // print bottom to top
		for iter.Next != nil {
			iter = iter.Next
		}
		for iter != nil {
			fmt.Println(iter.X)
			iter = iter.Prev
		}
	} else { // print top to bottom
		for iter != nil {
			fmt.Println(iter.X)
			iter = iter.Next
		}
	}
}