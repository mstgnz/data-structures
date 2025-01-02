package stack

import (
	"testing"
)

func TestLinkedListStack_New(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		tests := []struct {
			name string
			data int
			want []int
		}{
			{
				name: "create empty stack",
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
				stack := NewLinkedListStack[int](tt.data)
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("LinkedListStack() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("LinkedListStack() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})

	t.Run("string stack", func(t *testing.T) {
		tests := []struct {
			name string
			data string
			want []string
		}{
			{
				name: "create empty stack",
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
				stack := NewLinkedListStack[string](tt.data)
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("LinkedListStack() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("LinkedListStack() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func TestLinkedListStack_Push(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		tests := []struct {
			name string
			init int
			data []int
			want []int
		}{
			{
				name: "push to empty stack",
				init: 0,
				data: []int{1},
				want: []int{1, 0},
			},
			{
				name: "push multiple items",
				init: 0,
				data: []int{1, 2, 3},
				want: []int{3, 2, 1, 0},
			},
			{
				name: "push to non-empty stack",
				init: 1,
				data: []int{2, 3},
				want: []int{3, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stack := NewLinkedListStack[int](tt.init)
				for _, v := range tt.data {
					stack.Push(v)
				}
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Push() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("After Push() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})

	t.Run("string stack", func(t *testing.T) {
		tests := []struct {
			name string
			init string
			data []string
			want []string
		}{
			{
				name: "push to empty stack",
				init: "",
				data: []string{"a"},
				want: []string{"a", ""},
			},
			{
				name: "push multiple items",
				init: "",
				data: []string{"a", "b", "c"},
				want: []string{"c", "b", "a", ""},
			},
			{
				name: "push to non-empty stack",
				init: "a",
				data: []string{"b", "c"},
				want: []string{"c", "b", "a"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stack := NewLinkedListStack[string](tt.init)
				for _, v := range tt.data {
					stack.Push(v)
				}
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Push() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("After Push() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func TestLinkedListStack_Pop(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		tests := []struct {
			name string
			init int
			push []int
			pops int
			want []int
		}{
			{
				name: "pop from empty stack",
				init: 0,
				push: []int{},
				pops: 1,
				want: []int{0},
			},
			{
				name: "pop single item",
				init: 1,
				push: []int{},
				pops: 1,
				want: []int{0},
			},
			{
				name: "pop multiple items",
				init: 1,
				push: []int{2, 3},
				pops: 2,
				want: []int{1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stack := NewLinkedListStack[int](tt.init)
				for _, v := range tt.push {
					stack.Push(v)
				}
				for i := 0; i < tt.pops; i++ {
					stack.Pop()
				}
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Pop() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("After Pop() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})

	t.Run("string stack", func(t *testing.T) {
		tests := []struct {
			name string
			init string
			push []string
			pops int
			want []string
		}{
			{
				name: "pop from empty stack",
				init: "",
				push: []string{},
				pops: 1,
				want: []string{""},
			},
			{
				name: "pop single item",
				init: "a",
				push: []string{},
				pops: 1,
				want: []string{""},
			},
			{
				name: "pop multiple items",
				init: "a",
				push: []string{"b", "c"},
				pops: 2,
				want: []string{"a"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stack := NewLinkedListStack[string](tt.init)
				for _, v := range tt.push {
					stack.Push(v)
				}
				for i := 0; i < tt.pops; i++ {
					stack.Pop()
				}
				got := stack.List()
				if len(got) != len(tt.want) {
					t.Errorf("After Pop() = %v, want %v", got, tt.want)
					return
				}
				for i, v := range got {
					if v != tt.want[i] {
						t.Errorf("After Pop() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func BenchmarkLinkedListStack_Push(b *testing.B) {
	stack := NewLinkedListStack[int](0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkLinkedListStack_Pop(b *testing.B) {
	stack := NewLinkedListStack[int](0)
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			// Refill the stack when empty
			for j := 0; j < 1000; j++ {
				stack.Push(j)
			}
		}
		stack.Pop()
	}
}
