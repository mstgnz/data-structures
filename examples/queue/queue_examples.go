package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/queue"
)

// RunExamples demonstrates various queue implementations
func RunExamples() {
	// Example 1: Array-based Queue
	fmt.Println("Array Queue Example:")
	arrayQueue := queue.ArrayQueue()

	fmt.Println("Enqueuing elements:")
	arrayQueue.Enqueue(1)
	arrayQueue.Enqueue(2)
	arrayQueue.Enqueue(3)
	arrayQueue.Print()

	fmt.Println("\nDequeuing elements:")
	arrayQueue.Dequeue()
	arrayQueue.Print()

	fmt.Printf("Queue elements: %v\n\n", arrayQueue.List())

	// Example 2: LinkedList-based Queue
	fmt.Println("LinkedList Queue Example:")
	linkedQueue := queue.LinkedListQueue(0)

	fmt.Println("Enqueuing elements:")
	linkedQueue.Enqueue(10)
	linkedQueue.Enqueue(20)
	linkedQueue.Enqueue(30)
	linkedQueue.Print()

	fmt.Println("\nDequeuing elements:")
	linkedQueue.Dequeue()
	linkedQueue.Print()

	fmt.Printf("Queue elements: %v\n", linkedQueue.List())
}
