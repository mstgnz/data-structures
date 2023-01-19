package Test

import (
	"reflect"
	"testing"

	"github.com/mstgnz/data-structures/Queue"
)

func TestLinkedListQueue(t *testing.T) {
	queueLinkedList := Queue.LinkedListQueue(1)
	expect := []int{1}
	if got := queueLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("LinkedListQueue() = %v, want %v", got, expect)
	}
}

func Test_linkedListQueue_Dequeue(t *testing.T) {
	queueLinkedList := Queue.LinkedListQueue(1)
	expect := []int{5, 2}
	queueLinkedList.Enqueue(5)
	queueLinkedList.Dequeue()
	queueLinkedList.Enqueue(2)
	if got := queueLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Dequeue() = %v, want %v", got, expect)
	}
}

func Test_linkedListQueue_Enqueue(t *testing.T) {
	queueLinkedList := Queue.LinkedListQueue(1)
	expect := []int{3, 2}
	queueLinkedList.Enqueue(3)
	queueLinkedList.Enqueue(2)
	queueLinkedList.Dequeue()
	if got := queueLinkedList.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Enqueue() = %v, want %v", got, expect)
	}
}
