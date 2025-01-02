package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/linkedlist"
)

// RunExamples demonstrates various linked list implementations
func RunExamples() {
	// Example 1: Linear (Singly) Linked List
	fmt.Println("Linear Linked List Example:")
	list := linkedlist.NewLinear(0)

	fmt.Println("Adding elements to start:")
	list.AddToStart(1)
	list.AddToStart(2)
	list.AddToStart(3)
	list.Print()

	fmt.Println("\nAdding elements to end:")
	list.AddToEnd(4)
	list.AddToEnd(5)
	list.Print()

	fmt.Println("\nAdding elements sequentially:")
	list.AddToSequentially(2)
	list.AddToSequentially(6)
	list.Print()

	fmt.Printf("\nSearch for 4: %v\n\n", list.Search(4))

	// Example 2: Double Linked List
	fmt.Println("Double Linked List Example:")
	dlist := linkedlist.NewDouble(0)

	fmt.Println("Adding elements:")
	dlist.AddToStart(1)
	dlist.AddToEnd(2)
	dlist.AddToSequentially(3)

	fmt.Println("Forward traversal:")
	dlist.Print(false)
	fmt.Println("Reverse traversal:")
	dlist.Print(true)

	// Example 3: Circular Linked List
	fmt.Println("\nCircular Linked List Example:")
	clist := linkedlist.NewCircular(0)

	fmt.Println("Adding elements:")
	clist.AddToStart(1)
	clist.AddToEnd(2)
	clist.AddToSequentially(3)

	fmt.Println("List elements:")
	clist.Print()

	clist.AddToAfter(4, 2)

	fmt.Println("After adding 4 after 2:")
	clist.Print()
}
