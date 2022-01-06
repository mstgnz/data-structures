package Test

import (
	"data-structures/LinkedList"
	"reflect"
	"testing"
)

var circular LinkedList.ICircular = LinkedList.Circular(1)

func TestCircular(t *testing.T) {
	circular = LinkedList.Circular(1)
	expect := []int{1}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Circular() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToBetween(t *testing.T) {
	circular.AddToBetween(2,1)
	expect := []int{1,2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToBetween() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToEnd(t *testing.T) {
	circular.AddToEnd(2)
	expect := []int{1,2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToEnd() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToSequentially(t *testing.T) {
	circular.AddToSequentially(2)
	expect := []int{1,2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToSequentially() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToStart(t *testing.T) {
	circular.AddToStart(2)
	expect := []int{2,1}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToStart() = %v, want %v", got, expect)
	}
}

func Test_circular_Delete(t *testing.T) {
	circular.AddToStart(2)
	circular.Delete(1)
	expect := []int{2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}