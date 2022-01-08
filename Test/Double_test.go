package Test

import (
	"reflect"
	"testing"

	"data-structures/LinkedList"
)

func TestDouble(t *testing.T) {
	double := LinkedList.Double(1)
	expect := []int{1}
	if got := double.List(true); !reflect.DeepEqual(got, expect) {
		t.Errorf("Double() = %v, want %v", got, expect)
	}
}

func Test_double_AddToAfter(t *testing.T) {
	double := LinkedList.Double(1)
	double.AddToAfter(2,1)
	expect := []int{1,2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToAfter() = %v, want %v", got, expect)
	}
}

func Test_double_AddToEnd(t *testing.T) {
	double := LinkedList.Double(1)
	double.AddToEnd(2)
	expect := []int{1,2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToEnd() = %v, want %v", got, expect)
	}
}

func Test_double_AddToSequentially(t *testing.T) {
	double := LinkedList.Double(1)
	double.AddToSequentially(2)
	expect := []int{1,2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToSequentially() = %v, want %v", got, expect)
	}
}

func Test_double_AddToStart(t *testing.T) {
	double := LinkedList.Double(1)
	double.AddToStart(2)
	expect := []int{2,1}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToStart() = %v, want %v", got, expect)
	}
}

func Test_double_Delete(t *testing.T) {
	double := LinkedList.Double(1)
	double.AddToStart(2)
	double.Delete(1)
	expect := []int{2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}