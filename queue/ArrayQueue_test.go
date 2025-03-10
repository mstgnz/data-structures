package queue

import (
	"reflect"
	"testing"
)

func TestArrayQueue_New(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "create empty queue",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewArrayQueue[int]()
			got := queue.List()
			if len(got) != len(tt.want) {
				t.Errorf("NewArrayQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name string
			data []int
			want []int
		}{
			{
				name: "enqueue to empty queue",
				data: []int{1},
				want: []int{1},
			},
			{
				name: "enqueue multiple items",
				data: []int{1, 2, 3},
				want: []int{1, 2, 3},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewArrayQueue[int]()
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
	})

	t.Run("string queue", func(t *testing.T) {
		tests := []struct {
			name string
			data []string
			want []string
		}{
			{
				name: "enqueue to empty queue",
				data: []string{"a"},
				want: []string{"a"},
			},
			{
				name: "enqueue multiple items",
				data: []string{"a", "b", "c"},
				want: []string{"a", "b", "c"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewArrayQueue[string]()
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
	})
}

func TestArrayQueue_Dequeue(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name     string
			init     []int
			dequeues int
			want     []int
		}{
			{
				name:     "dequeue from empty queue",
				init:     []int{},
				dequeues: 1,
				want:     []int{},
			},
			{
				name:     "dequeue single item",
				init:     []int{1},
				dequeues: 1,
				want:     []int{},
			},
			{
				name:     "dequeue multiple items",
				init:     []int{1, 2, 3},
				dequeues: 2,
				want:     []int{3},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewArrayQueue[int]()
				for _, v := range tt.init {
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
	})

	t.Run("string queue", func(t *testing.T) {
		tests := []struct {
			name     string
			init     []string
			dequeues int
			want     []string
		}{
			{
				name:     "dequeue from empty queue",
				init:     []string{},
				dequeues: 1,
				want:     []string{},
			},
			{
				name:     "dequeue single item",
				init:     []string{"a"},
				dequeues: 1,
				want:     []string{},
			},
			{
				name:     "dequeue multiple items",
				init:     []string{"a", "b", "c"},
				dequeues: 2,
				want:     []string{"c"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewArrayQueue[string]()
				for _, v := range tt.init {
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
	})
}

func TestArrayQueue_Print(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
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
				queue := NewArrayQueue[int]()
				for _, v := range tt.data {
					queue.Enqueue(v)
				}
				queue.Print()
			})
		}
	})

	t.Run("string queue", func(t *testing.T) {
		tests := []struct {
			name string
			data []string
		}{
			{
				name: "print empty queue",
				data: []string{},
			},
			{
				name: "print single item",
				data: []string{"a"},
			},
			{
				name: "print multiple items",
				data: []string{"a", "b", "c"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewArrayQueue[string]()
				for _, v := range tt.data {
					queue.Enqueue(v)
				}
				queue.Print()
			})
		}
	})
}

func TestArrayQueue_ReSort(t *testing.T) {
	queue := NewArrayQueue[int]()
	// Fill the queue
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}
	// Dequeue some elements to create gaps
	for i := 0; i < 5; i++ {
		queue.Dequeue()
	}
	// Enqueue more elements to trigger reSort
	for i := 10; i < 15; i++ {
		queue.Enqueue(i)
	}

	// Verify the queue is still correct after reSort
	expected := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	got := queue.List()
	if len(got) != len(expected) {
		t.Errorf("After reSort() got %v, want %v", got, expected)
		return
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("After reSort() got %v, want %v", got, expected)
		}
	}
}

func TestArrayQueue_reSort(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name     string
			setup    func() *ArrayQueue[int]
			expected []int
		}{
			{
				name: "resort_empty_queue",
				setup: func() *ArrayQueue[int] {
					queue := NewArrayQueue[int]()
					return queue
				},
				expected: []int{},
			},
			{
				name: "resort_single_element",
				setup: func() *ArrayQueue[int] {
					queue := NewArrayQueue[int]()
					queue.Enqueue(1)
					queue.Dequeue()
					return queue
				},
				expected: []int{},
			},
			{
				name: "resort_multiple_elements",
				setup: func() *ArrayQueue[int] {
					queue := NewArrayQueue[int]()
					queue.Enqueue(1)
					queue.Enqueue(2)
					queue.Enqueue(3)
					queue.Dequeue()
					queue.Dequeue()
					queue.Enqueue(4) // This will trigger reSort
					return queue
				},
				expected: []int{3, 4},
			},
			{
				name: "resort_full_queue",
				setup: func() *ArrayQueue[int] {
					queue := NewArrayQueue[int]()
					for i := 0; i < 10; i++ {
						queue.Enqueue(i)
					}
					for i := 0; i < 5; i++ {
						queue.Dequeue()
					}
					queue.Enqueue(10) // This will trigger reSort
					return queue
				},
				expected: []int{5, 6, 7, 8, 9, 10},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := tt.setup()
				result := queue.List()
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("reSort() got %v, want %v", result, tt.expected)
				}
			})
		}
	})
}

func BenchmarkArrayQueue_Enqueue(b *testing.B) {
	queue := NewArrayQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkArrayQueue_Dequeue(b *testing.B) {
	queue := NewArrayQueue[int]()
	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}
