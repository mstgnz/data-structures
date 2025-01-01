package tree

import (
	"fmt"
	"sync"
)

type IBinaryTree interface {
	Insert(data int)
	Search(data int)
	Exists(data int) bool
	Delete(data int)
	Max() int
	Min() int
	Print(pType string)
	List(pType string) []int
}

type binaryTree struct {
	X     int
	Left  *binaryTree
	Right *binaryTree
	mutex sync.RWMutex
}

func BinaryTree(data int) IBinaryTree {
	return &binaryTree{X: data, Left: nil, Right: nil, mutex: sync.RWMutex{}}
}

// Insert adds data to the tree
func (tree *binaryTree) Insert(data int) {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	recursiveInsert(data, tree)
}

// Search prints whether data exists in the tree
func (tree *binaryTree) Search(data int) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	if recursiveSearch(data, tree) {
		fmt.Printf("%v: available in the tree\n", data)
	} else {
		fmt.Printf("%v: value not found\n", data)
	}
}

// Exists returns true if data exists in the tree
func (tree *binaryTree) Exists(data int) bool {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	return recursiveSearch(data, tree)
}

// Max returns the maximum value in the tree
func (tree *binaryTree) Max() int {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	iter := tree
	for iter.Right != nil {
		iter = iter.Right
	}
	return iter.X
}

// Min returns the minimum value in the tree
func (tree *binaryTree) Min() int {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	iter := tree
	for iter.Left != nil {
		iter = iter.Left
	}
	return iter.X
}

// Delete removes data from the tree
func (tree *binaryTree) Delete(data int) {
	tree.mutex.Lock()
	defer tree.mutex.Unlock()
	recursiveDelete(data, tree)
}

// List returns a slice of values in the specified traversal order
// Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN
func (tree *binaryTree) List(pType string) []int {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	var list []int
	switch pType {
	case "NLR":
	case "NRL":
		prefixPrint(tree, pType, &list)
	case "LRN":
	case "RLN":
		postfixPrint(tree, pType, &list)
	default:
		infixPrint(tree, pType, &list)
	}
	return list
}

// Print displays the tree values in the specified traversal order
func (tree *binaryTree) Print(pType string) {
	tree.mutex.RLock()
	defer tree.mutex.RUnlock()
	fmt.Print("print : ")
	for _, val := range tree.List(pType) {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// Infix: LNR-RNL
func infixPrint(tree *binaryTree, pType string, list *[]int) {
	if tree == nil {
		return
	}
	if pType == "RNL" {
		infixPrint(tree.Right, pType, list)
		*list = append(*list, tree.X)
		infixPrint(tree.Left, pType, list)
	} else {
		infixPrint(tree.Left, pType, list)
		*list = append(*list, tree.X)
		infixPrint(tree.Right, pType, list)
	}
}

// Prefix: NLR-NRL
func prefixPrint(tree *binaryTree, pType string, list *[]int) {
	if tree == nil {
		return
	}
	if pType == "NRL" {
		*list = append(*list, tree.X)
		infixPrint(tree.Right, pType, list)
		infixPrint(tree.Left, pType, list)
	} else {
		*list = append(*list, tree.X)
		infixPrint(tree.Left, pType, list)
		infixPrint(tree.Right, pType, list)
	}
}

// Postfix: LRN, RLN
func postfixPrint(tree *binaryTree, pType string, list *[]int) {
	if tree == nil {
		return
	}
	if pType == "RLN" {
		infixPrint(tree.Right, pType, list)
		infixPrint(tree.Left, pType, list)
		*list = append(*list, tree.X)
	} else {
		infixPrint(tree.Left, pType, list)
		infixPrint(tree.Right, pType, list)
		*list = append(*list, tree.X)
	}
}

// recursive insert
func recursiveInsert(data int, tree *binaryTree) *binaryTree {
	if tree != nil {
		if tree.X < data {
			tree.Right = recursiveInsert(data, tree.Right)
		} else {
			tree.Left = recursiveInsert(data, tree.Left)
		}
	} else {
		tree = &binaryTree{X: data, Left: nil, Right: nil}
	}
	return tree
}

// recursive search
func recursiveSearch(data int, tree *binaryTree) bool {
	if tree == nil {
		return false
	}
	if tree.X == data {
		return true
	}
	if recursiveSearch(data, tree.Left) {
		return true
	}
	if recursiveSearch(data, tree.Right) {
		return true
	}
	return false
}

// recursive delete
func recursiveDelete(data int, tree *binaryTree) *binaryTree {
	if tree == nil {
		return nil
	}
	if tree.X == data {
		if tree.Left == nil && tree.Right == nil {
			return nil
		}
		if tree.Right != nil {
			tree.X = min(tree.Right)
			tree.Right = recursiveDelete(min(tree.Right), tree.Right)
		} else {
			tree.X = max(tree.Left)
			tree.Left = recursiveDelete(max(tree.Left), tree.Left)
		}
	} else {
		if tree.X < data {
			tree.Right = recursiveDelete(data, tree.Right)
		} else {
			tree.Left = recursiveDelete(data, tree.Left)
		}
	}
	return tree
}

// Max
func max(tree *binaryTree) int {
	iter := tree
	for iter.Right != nil {
		iter = iter.Right
	}
	return iter.X
}

// Min
func min(tree *binaryTree) int {
	iter := tree
	for iter.Left != nil {
		iter = iter.Left
	}
	return iter.X
}
