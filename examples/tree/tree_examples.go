package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/tree"
)

// RunExamples demonstrates various tree data structures and algorithms
func RunExamples() {
	// Example 1: Binary Tree
	fmt.Println("Binary Tree Example:")
	binaryTree := tree.BinaryTree(50)
	values := []int{30, 70, 20, 40, 60, 80}
	for _, value := range values {
		binaryTree.Insert(value)
	}

	fmt.Print("Infix (LNR): ")
	binaryTree.Print("LNR")
	fmt.Print("Prefix (NLR): ")
	binaryTree.Print("NLR")
	fmt.Print("Postfix (LRN): ")
	binaryTree.Print("LRN")
	fmt.Printf("Min: %d, Max: %d\n\n", binaryTree.Min(), binaryTree.Max())

	// Example 2: AVL Tree
	fmt.Println("AVL Tree Example:")
	avlTree := tree.NewAVLTree()
	avlValues := []int{10, 20, 30, 40, 50, 25}
	for _, value := range avlValues {
		avlTree.Insert(value)
	}

	var result []int
	avlTree.InOrderTraversal(&result)
	fmt.Printf("AVL Tree (Inorder): %v\n", result)
	fmt.Printf("Search for 25: %v\n\n", avlTree.Search(25))

	// Example 3: Red-Black Tree
	fmt.Println("Red-Black Tree Example:")
	rbTree := tree.NewRedBlackTree()
	rbValues := []int{7, 3, 18, 10, 22, 8, 11, 26}
	for _, value := range rbValues {
		rbTree.Insert(value)
	}

	var rbResult []int
	rbTree.InOrderTraversal(&rbResult)
	fmt.Printf("RB Tree (Inorder): %v\n", rbResult)
	fmt.Printf("Search for 18: %v\n\n", rbTree.Search(18))

	// Example 4: B-Tree
	fmt.Println("B-Tree Example:")
	bTree := tree.NewBTree(3) // Minimum degree = 3
	bValues := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, value := range bValues {
		bTree.Insert(value)
	}

	fmt.Printf("B-Tree (Inorder): %v\n", bTree.GetInOrder())
	fmt.Printf("Search for 17: %v\n", bTree.Search(17))
}
