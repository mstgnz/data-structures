package Tree

import (
	"reflect"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		rbt := NewRedBlackTree()

		// Test insertion
		values := []int{7, 3, 18, 10, 22, 8, 11, 26, 2, 6}
		for _, v := range values {
			rbt.Insert(v)
		}

		// Test search
		for _, v := range values {
			if !rbt.Search(v) {
				t.Errorf("Value %d should be found in the tree", v)
			}
		}

		// Test non-existent value
		if rbt.Search(100) {
			t.Error("Value 100 should not be found in the tree")
		}

		// Test inorder traversal
		var result []int
		rbt.InOrderTraversal(&result)
		expected := []int{2, 3, 6, 7, 8, 10, 11, 18, 22, 26}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected inorder traversal %v, got %v", expected, result)
		}

		// Test Red-Black properties
		if rbt.Root.Color != BLACK {
			t.Error("Root should be black")
		}

		// Test that every path from root to leaf has same number of black nodes
		blackHeight := getBlackHeight(rbt.Root, rbt.NIL)
		if !validateBlackHeight(rbt.Root, rbt.NIL, blackHeight, 0) {
			t.Error("Black height property violated")
		}

		// Test that no red node has red child
		if !validateRedProperty(rbt.Root, rbt.NIL) {
			t.Error("Red property violated")
		}
	})
}

// Helper function to get black height of the tree
func getBlackHeight(node *RBNode, nil_node *RBNode) int {
	if node == nil_node {
		return 0
	}
	height := getBlackHeight(node.Left, nil_node)
	if node.Color == BLACK {
		height++
	}
	return height
}

// Helper function to validate black height property
func validateBlackHeight(node *RBNode, nil_node *RBNode, expectedHeight int, currentHeight int) bool {
	if node == nil_node {
		return currentHeight == expectedHeight
	}
	if node.Color == BLACK {
		currentHeight++
	}
	return validateBlackHeight(node.Left, nil_node, expectedHeight, currentHeight) &&
		validateBlackHeight(node.Right, nil_node, expectedHeight, currentHeight)
}

// Helper function to validate red property
func validateRedProperty(node *RBNode, nil_node *RBNode) bool {
	if node == nil_node {
		return true
	}
	if node.Color == RED {
		if node.Left.Color == RED || node.Right.Color == RED {
			return false
		}
	}
	return validateRedProperty(node.Left, nil_node) && validateRedProperty(node.Right, nil_node)
}
