package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/stack"
)

// RunExamples demonstrates various stack implementations
func RunExamples() {
	// Example 1: Array-based Stack
	fmt.Println("Array Stack Example:")
	arrayStack := stack.ArrayStack()

	fmt.Println("Pushing elements:")
	arrayStack.Push(1)
	arrayStack.Push(2)
	arrayStack.Push(3)
	arrayStack.Print()

	fmt.Println("\nPopping elements:")
	arrayStack.Pop()
	arrayStack.Print()

	fmt.Printf("\nStack is empty: %v\n", arrayStack.IsEmpty())
	fmt.Printf("Stack elements: %v\n\n", arrayStack.List())

	// Example 2: LinkedList-based Stack
	fmt.Println("LinkedList Stack Example:")
	linkedStack := stack.LinkedListStack(0)

	fmt.Println("Pushing elements:")
	linkedStack.Push(10)
	linkedStack.Push(20)
	linkedStack.Push(30)
	linkedStack.Print()

	fmt.Println("\nPopping elements:")
	linkedStack.Pop()
	linkedStack.Print()

	fmt.Printf("\nStack is empty: %v\n", linkedStack.IsEmpty())
	fmt.Printf("Stack elements: %v\n", linkedStack.List())
}
