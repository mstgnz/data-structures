package Test

import (
	"reflect"
	"testing"

	"data-structures/Stack"
)

func TestArrayStack(t *testing.T) {
	stackArray := Stack.ArrayStack()
	var expect []int
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("ArrayStack() = %v, want %v", got, expect)
	}
}

func Test_arrayStack_Pop(t *testing.T) {
	stackArray := Stack.ArrayStack()
	expect := []int{1}
	stackArray.Push(1)
	stackArray.Push(2)
	stackArray.Pop()
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Pop() = %v, want %v", got, expect)
	}
}

func Test_arrayStack_Push(t *testing.T) {
	stackArray := Stack.ArrayStack()
	expect := []int{5,9}
	stackArray.Push(5)
	stackArray.Push(3)
	stackArray.Pop()
	stackArray.Push(9)
	if got := stackArray.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Push() = %v, want %v", got, expect)
	}
}
