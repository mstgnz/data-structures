package queue

import (
	"testing"
)

func TestLinkedListQueue_New(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name string
			data int
			want []int
		}{
			{
				name: "create empty queue",
				data: 0,
				want: []int{0},
			},
			{
				name: "create with value",
				data: 1,
				want: []int{1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewLinkedListQueue[int](tt.data)
				got := queue.List()
				if len(got) != len(tt.want) {
					t.Errorf("LinkedListQueue() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("LinkedListQueue() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})

	t.Run("string queue", func(t *testing.T) {
		tests := []struct {
			name string
			data string
			want []string
		}{
			{
				name: "create empty queue",
				data: "",
				want: []string{""},
			},
			{
				name: "create with value",
				data: "a",
				want: []string{"a"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewLinkedListQueue[string](tt.data)
				got := queue.List()
				if len(got) != len(tt.want) {
					t.Errorf("LinkedListQueue() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("LinkedListQueue() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func TestLinkedListQueue_Enqueue(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name string
			init int
			data []int
			want []int
		}{
			{
				name: "enqueue to empty queue",
				init: 0,
				data: []int{1},
				want: []int{0, 1},
			},
			{
				name: "enqueue multiple items",
				init: 0,
				data: []int{1, 2, 3},
				want: []int{0, 1, 2, 3},
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
				queue := NewLinkedListQueue[int](tt.init)
				for _, v := range tt.data {
					queue.Enqueue(v)
				}
				got := queue.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Enqueue() = %v, want %v", got, tt.want)
					return
				}
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
			init string
			data []string
			want []string
		}{
			{
				name: "enqueue to empty queue",
				init: "",
				data: []string{"a"},
				want: []string{"", "a"},
			},
			{
				name: "enqueue multiple items",
				init: "",
				data: []string{"a", "b", "c"},
				want: []string{"", "a", "b", "c"},
			},
			{
				name: "enqueue to non-empty queue",
				init: "a",
				data: []string{"b", "c"},
				want: []string{"a", "b", "c"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewLinkedListQueue[string](tt.init)
				for _, v := range tt.data {
					queue.Enqueue(v)
				}
				got := queue.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Enqueue() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("After Enqueue() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	t.Run("integer queue", func(t *testing.T) {
		tests := []struct {
			name     string
			init     int
			enqueue  []int
			dequeues int
			want     []int
		}{
			{
				name:     "dequeue from empty queue",
				init:     0,
				enqueue:  []int{},
				dequeues: 1,
				want:     []int{0},
			},
			{
				name:     "dequeue single item",
				init:     1,
				enqueue:  []int{},
				dequeues: 1,
				want:     []int{0},
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
				queue := NewLinkedListQueue[int](tt.init)
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
	})

	t.Run("string queue", func(t *testing.T) {
		tests := []struct {
			name     string
			init     string
			enqueue  []string
			dequeues int
			want     []string
		}{
			{
				name:     "dequeue from empty queue",
				init:     "",
				enqueue:  []string{},
				dequeues: 1,
				want:     []string{""},
			},
			{
				name:     "dequeue single item",
				init:     "a",
				enqueue:  []string{},
				dequeues: 1,
				want:     []string{""},
			},
			{
				name:     "dequeue multiple items",
				init:     "a",
				enqueue:  []string{"b", "c"},
				dequeues: 2,
				want:     []string{"c"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				queue := NewLinkedListQueue[string](tt.init)
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
	})
}

func TestLinkedListQueue_Print(t *testing.T) {
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
				queue := NewLinkedListQueue[int](0)
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
				queue := NewLinkedListQueue[string]("")
				for _, v := range tt.data {
					queue.Enqueue(v)
				}
				queue.Print()
			})
		}
	})
}

func BenchmarkLinkedListQueue_Enqueue(b *testing.B) {
	queue := NewLinkedListQueue[int](0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkLinkedListQueue_Dequeue(b *testing.B) {
	queue := NewLinkedListQueue[int](0)
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
