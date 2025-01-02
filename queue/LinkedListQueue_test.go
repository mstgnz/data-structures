package queue

import (
	"testing"
)

func TestLinkedListQueue_New(t *testing.T) {
	tests := []struct {
		name string
		data int
		want []int
	}{
		{
			name: "create empty queue",
			data: -1,
			want: []int{-1},
		},
		{
			name: "create with value",
			data: 1,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewLinkedListQueue(tt.data)
			got := queue.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("LinkedListQueue() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinkedListQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name string
		init int
		data []int
		want []int
	}{
		{
			name: "enqueue to empty queue",
			init: -1,
			data: []int{1},
			want: []int{1},
		},
		{
			name: "enqueue multiple items",
			init: -1,
			data: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "enqueue to non-empty queue",
			init: 1,
			data: []int{2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewLinkedListQueue(tt.init)
			for _, v := range tt.data {
				queue.Enqueue(v)
			}
			got := queue.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After Enqueue() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		enqueue  []int
		dequeues int
		want     []int
	}{
		{
			name:     "dequeue from empty queue",
			init:     -1,
			enqueue:  []int{},
			dequeues: 1,
			want:     []int{-1},
		},
		{
			name:     "dequeue single item",
			init:     1,
			enqueue:  []int{},
			dequeues: 1,
			want:     []int{-1},
		},
		{
			name:     "dequeue multiple items",
			init:     1,
			enqueue:  []int{2, 3},
			dequeues: 2,
			want:     []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewLinkedListQueue(tt.init)
			for _, v := range tt.enqueue {
				queue.Enqueue(v)
			}
			for i := 0; i < tt.dequeues; i++ {
				queue.Dequeue()
			}
			got := queue.List()
			if len(got) != len(tt.want) {
				t.Errorf("After Dequeue() = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After Dequeue() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinkedListQueue_Print(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "print empty queue",
			data: []int{},
		},
		{
			name: "print single item",
			data: []int{1},
		},
		{
			name: "print multiple items",
			data: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewLinkedListQueue(0)
			for _, v := range tt.data {
				queue.Enqueue(v)
			}
			queue.Print()
		})
	}
}

func BenchmarkLinkedListQueue_Enqueue(b *testing.B) {
	queue := NewLinkedListQueue(-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkLinkedListQueue_Dequeue(b *testing.B) {
	queue := NewLinkedListQueue(-1)
	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			// Refill the queue when empty
			for j := 0; j < 1000; j++ {
				queue.Enqueue(j)
			}
		}
		queue.Dequeue()
	}
}
