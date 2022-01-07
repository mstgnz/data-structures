package Test

import (
	"data-structures/Stack"
	"reflect"
	"testing"
)

var stackArray Stack.IArrayStack = Stack.ArrayStack()

func TestArrayStack(t *testing.T) {
	var expect []int
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("ArrayStack() = %v, want %v", got, expect)
	}
}

func Test_arrayStack_Pop(t *testing.T) {
	expect := []int{1}
	stackArray.Push(1)
	stackArray.Push(2)
	stackArray.Pop()
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Pop() = %v, want %v", got, expect)
	}
}

func Test_arrayStack_Push(t *testing.T) {
	expect := []int{5,9}
	stackArray.Push(5)
	stackArray.Push(3)
	stackArray.Pop()
	stackArray.Push(9)
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Push() = %v, want %v", got, expect)
	}
}
