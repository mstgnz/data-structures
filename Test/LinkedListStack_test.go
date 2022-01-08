package Test

import (
	"reflect"
	"testing"

	"data-structures/Stack"
)

func TestLinkedListStack(t *testing.T) {
	stackLinkedList := Stack.LinkedListStack(1)
	expect :=  []int{1}
	if got := stackLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("LinkedListStack() = %v, want %v", got, expect)
	}
}

func Test_linkedListStack_Pop(t *testing.T) {
	stackLinkedList := Stack.LinkedListStack(1)
	expect :=  []int{-1}
	stackLinkedList.Pop()
	if got := stackLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Pop() = %v, want %v", got, expect)
	}
}

func Test_linkedListStack_Push(t *testing.T) {
	stackLinkedList := Stack.LinkedListStack(1)
	expect :=  []int{3,1}
	stackLinkedList.Pop()
	stackLinkedList.Push(3)
	stackLinkedList.Push(1)
	if got := stackLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Push() = %v, want %v", got, expect)
	}
}
