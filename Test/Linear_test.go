package Test

import (
	"reflect"
	"testing"

	"github.com/mstgnz/data-structures/LinkedList"
)

func TestLinear(t *testing.T) {
	linear := LinkedList.Linear(1)
	expect := []int{1}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Linear() = %v, want %v", got, expect)
	}
}

func Test_linear_AddToAfter(t *testing.T) {
	linear := LinkedList.Linear(1)
	linear.AddToAfter(2, 1)
	expect := []int{1, 2}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToAfter() = %v, want %v", got, expect)
	}
}

func Test_linear_AddToEnd(t *testing.T) {
	linear := LinkedList.Linear(1)
	linear.AddToEnd(2)
	linear.AddToAfter(3, 1)
	expect := []int{1, 3, 2}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToEnd() = %v, want %v", got, expect)
	}
}

func Test_linear_AddToSequentially(t *testing.T) {
	linear := LinkedList.Linear(1)
	linear.AddToSequentially(2)
	expect := []int{1, 2}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToSequentially() = %v, want %v", got, expect)
	}
}

func Test_linear_AddToStart(t *testing.T) {
	linear := LinkedList.Linear(1)
	linear.AddToStart(2)
	expect := []int{2, 1}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToStart() = %v, want %v", got, expect)
	}
}

func Test_linear_Delete(t *testing.T) {
	linear := LinkedList.Linear(1)
	linear.AddToStart(2)
	linear.Delete(1)
	expect := []int{2}
	if got := linear.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}
