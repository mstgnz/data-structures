package Tree

import "fmt"

type IBinaryTree interface{
	Insert(data int)
	Search(data int)
	Exists(data int)bool
	Delete(data int)
	Max()int
	Min()int
	Print(pType string)
	List(pType string) []int
}

type binaryTree struct {
	X int
	Left *binaryTree
	Right *binaryTree
}

func BinaryTree(data int) IBinaryTree{
	return &binaryTree{data, nil, nil}
}

// Insert Add to data
func (tree *binaryTree) Insert(data int) {
	recursiveInsert(data, tree)
}

// Search print
func (tree *binaryTree) Search(data int) {
	if recursiveSearch(data, tree){
		fmt.Printf("%v: available in the tree\n", data)
	}else{
		fmt.Printf("%v: value not found\n", data)
	}
}

// Exists true or false
func (tree *binaryTree) Exists(data int) bool {
	return recursiveSearch(data, tree)
}

// Max find value
func (tree *binaryTree) Max() int {
	iter := tree
	for iter.Right != nil {
		iter = iter.Right
	}
	return iter.X
}

// Min find value
func (tree *binaryTree) Min() int {
	iter := tree
	for iter.Left != nil {
		iter = iter.Left
	}
	return iter.X
}

// Delete Remove to data
func (tree *binaryTree) Delete(data int) {
	tree = recursiveDelete(data, tree)
}

// List Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN
func (tree *binaryTree) List(pType string) []int {
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

// Print Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN
func (tree *binaryTree) Print(pType string) {
	fmt.Print("print : ")
	for _, val := range tree.List(pType) {
		fmt.Print(val," ")
	}
	fmt.Println()
}

// Infix: LNR-RNL
func infixPrint(tree *binaryTree, pType string, list *[]int){
	if tree == nil{
		return
	}
	if pType == "RNL"{
		infixPrint(tree.Right, pType, list)
		*list = append(*list, tree.X)
		infixPrint(tree.Left, pType, list)
	}else{
		infixPrint(tree.Left, pType, list)
		*list = append(*list, tree.X)
		infixPrint(tree.Right, pType, list)
	}
}

// Prefix: NLR-NRL
func prefixPrint(tree *binaryTree, pType string, list *[]int){
	if tree == nil{
		return
	}
	if pType == "NRL"{
		*list = append(*list, tree.X)
		infixPrint(tree.Right, pType, list)
		infixPrint(tree.Left, pType, list)
	}else{
		*list = append(*list, tree.X)
		infixPrint(tree.Left, pType, list)
		infixPrint(tree.Right, pType, list)
	}
}

// Postfix: LRN, RLN
func postfixPrint(tree *binaryTree, pType string, list *[]int){
	if tree == nil{
		return
	}
	if pType == "RLN"{
		infixPrint(tree.Right, pType, list)
		infixPrint(tree.Left, pType, list)
		*list = append(*list, tree.X)
	}else{
		infixPrint(tree.Left, pType, list)
		infixPrint(tree.Right, pType, list)
		*list = append(*list, tree.X)
	}
}

// recursive insert
func recursiveInsert(data int, tree *binaryTree) *binaryTree{
	if tree != nil{
		if tree.X < data{
			tree.Right = recursiveInsert(data, tree.Right)
		}else{
			tree.Left = recursiveInsert(data, tree.Left)
		}
	}else{
		tree = &binaryTree{X: data, Left: nil, Right: nil}
	}
	return tree
}

// recursive search
func recursiveSearch(data int, tree *binaryTree) bool{
	if tree == nil{
		return false
	}
	if tree.X == data{
		return true
	}
	if recursiveSearch(data, tree.Left){
		return true
	}
	if recursiveSearch(data, tree.Right){
		return true
	}
	return false
}

// recursive delete
func recursiveDelete(data int, tree *binaryTree) *binaryTree{
	if tree == nil{
		return nil
	}
	if tree.X == data{
		if tree.Left == nil && tree.Right == nil {
			return nil
		}
		if tree.Right != nil {
			tree.X = min(tree.Right)
			tree.Right = recursiveDelete(min(tree.Right), tree.Right)
		}else{
			tree.X = max(tree.Left)
			tree.Left = recursiveDelete(max(tree.Left), tree.Left)
		}
	}else{
		if tree.X < data{
			tree.Right = recursiveDelete(data, tree.Right)
		}else{
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