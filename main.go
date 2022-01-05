package main

import (
	"data-structures/LinkedList"
	"data-structures/Queue"
	"data-structures/Stack"
	"data-structures/Tree"
	"fmt"
)

func main(){
	// Linear Linked List, Each node keeps the reference of the next node object.
	linearLinkedList()

	// Circular Linked List holds the reference of the last node root object.
	// circularLinkedList()

	// Double Linked List, Each node keeps the reference of the next and previous node object. Bidirectional linking is provided.
	// doubleLinkedList()

	// Array Stack, Last in first out
	// stackArray()

	// Linked List Stack, Last in first out
	// stackLinkedList()

	// Array Queue, First in first out
	// queueArray()

	// Linked List Queue, First in first out
	// queueLinkedList()

	// Binary Tree
	// binaryTree()
}

// Linear Linked List
func linearLinkedList() {
	linear := LinkedList.Linear(1)
	linear.AddToEnd(33)
	linear.Delete(1)
	linear.AddToEnd(21)
	linear.AddToEnd(33)
	linear.AddToEnd(44)
	linear.AddToStart(3)
	linear.AddToSequentially(5)
	linear.AddToSequentially(35)
	linear.AddToSequentially(55)
	linear.AddToBetween(12, 21)
	linear.AddToBetween(66, 35)
	linear.Delete(23)
	linear.Delete(3)
	linear.Print()
}

// Circular Linked List
func circularLinkedList() {
	circular := LinkedList.Circular(1)
	circular.AddToSequentially(11)
	circular.AddToBetween(66, 11)
	circular.AddToSequentially(22)
	circular.AddToStart(55)
	circular.Delete(11)
	circular.Print()
}

// Double Linked List
func doubleLinkedList() {
	double := LinkedList.Double(1)
	double.AddToStart(3)
	double.AddToEnd(9)
	double.AddToEnd(12)
	double.AddToEnd(15)
	double.Delete(3)
	double.Delete(12)
	double.AddToBetween(20, 5)
	double.AddToSequentially(4)
	double.AddToSequentially(11)
	double.AddToStart(1)
	double.AddToStart(2)
	double.Print(false)
	fmt.Println("-----")
	double.Print(true)
}

// Array Stack
func stackArray(){
	myArrayStack := Stack.ArrayStack()
	//myArrayStack.Constructor()
	myArrayStack.Push(33)
	myArrayStack.Push(55)
	myArrayStack.Print()
	myArrayStack.Push(65)
	myArrayStack.Print()
	myArrayStack.Push(76)
	myArrayStack.Print()
	myArrayStack.Push(86)
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
}

// LinkedList Stack
func stackLinkedList(){
	myArrayStack := Stack.LinkedListStack(22)
	myArrayStack.Push(33)
	myArrayStack.Push(55)
	myArrayStack.Push(66)
	myArrayStack.Print()
	myArrayStack.Push(65)
	myArrayStack.Print()
	myArrayStack.Push(76)
	myArrayStack.Print()
	myArrayStack.Push(86)
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Pop()
	myArrayStack.Pop()
	myArrayStack.Pop()
	myArrayStack.Pop()
	myArrayStack.Print()
	myArrayStack.Push(55)
	myArrayStack.Push(66)
	myArrayStack.Print()
}

// Array Queue
func queueArray(){
	myArrayStack := Queue.ArrayQueue()
	myArrayStack.Enqueue(33)
	myArrayStack.Enqueue(55)
	myArrayStack.Enqueue(66)
	myArrayStack.Enqueue(77)
	myArrayStack.Dequeue()
	myArrayStack.Enqueue(33)
	myArrayStack.Print()
	myArrayStack.Enqueue(65)
	myArrayStack.Print()
	myArrayStack.Enqueue(76)
	myArrayStack.Print()
	myArrayStack.Enqueue(86)
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Enqueue(33)
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
}

// LinkedList Queue
func queueLinkedList(){
	myArrayStack := Queue.LinkedListQueue(22)
	myArrayStack.Enqueue(33)
	myArrayStack.Enqueue(55)
	myArrayStack.Enqueue(66)
	myArrayStack.Print()
	myArrayStack.Enqueue(65)
	myArrayStack.Print()
	myArrayStack.Enqueue(76)
	myArrayStack.Print()
	myArrayStack.Enqueue(86)
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Dequeue()
	myArrayStack.Dequeue()
	myArrayStack.Dequeue()
	myArrayStack.Dequeue()
	myArrayStack.Print()
	myArrayStack.Enqueue(55)
	myArrayStack.Enqueue(66)
	myArrayStack.Print()
}

// Binary Tree
func binaryTree(){
	myTree := Tree.BinaryTree(56)
	myTree.Insert(200)
	myTree.Insert(26)
	myTree.Insert(190)
	myTree.Insert(213)
	myTree.Insert(18)
	myTree.Insert(28)
	myTree.Insert(12)
	myTree.Insert(24)
	myTree.Insert(27)
	myTree.Print("NRL") // Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN
	myTree.Search(100)
	myTree.Search(24)
	fmt.Println(myTree.Exists(100))
	fmt.Println(myTree.Exists(24))
	fmt.Printf("Max value: %v\n", myTree.Max())
	fmt.Printf("Min value: %v\n", myTree.Min())
	myTree.Delete(56)
	myTree.Print("NRL")
}