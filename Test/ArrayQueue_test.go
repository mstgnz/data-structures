package Test

import (
	"reflect"
	"testing"

	"github.com/mstgnz/data-structures/Queue"
)

func TestArrayQueue(t *testing.T) {
	arrayQueue := Queue.ArrayQueue()
	var expect []int
	if got := arrayQueue.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("ArrayQueue() = %v, want %v", got, expect)
	}
}

func Test_arrayQueue_Dequeue(t *testing.T) {
	arrayQueue := Queue.ArrayQueue()
	var expect []int
	arrayQueue.Enqueue(1)
	arrayQueue.Dequeue()
	if got := arrayQueue.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Dequeue() = %v, want %v", got, expect)
	}
}

func Test_arrayQueue_Enqueue(t *testing.T) {
	arrayQueue := Queue.ArrayQueue()
	expect := []int{2}
	arrayQueue.Enqueue(1)
	arrayQueue.Dequeue()
	arrayQueue.Enqueue(2)
	if got := arrayQueue.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Enqueue() = %v, want %v", got, expect)
	}
}
