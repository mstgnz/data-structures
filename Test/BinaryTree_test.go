package Test

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/mstgnz/data-structures/Tree"
)

// tree.List(pType string) -> pType -> Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN

func TestBinaryTree(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := []int{1}
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("BinaryTree() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Delete(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := []int{1, 3} // for NRL
	tree.Insert(2)
	tree.Insert(3)
	tree.Delete(2)
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Exists(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := false
	got := tree.Exists(13)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Exists() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Insert(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := []int{1, 2, 3} // for NRL
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("Insert() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Max(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := 3
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.Max(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Max() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Min(t *testing.T) {
	tree := Tree.BinaryTree(1)
	expect := 1
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.Min(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Min() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Search(t *testing.T) {
	/*expect := "1: available in the tree"
	got := captureOutput(func() {
		tree.Search(1)
	})
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Search() = %v, want %v", got, expect)
	}*/
}

// captureOutput print
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
