package stack

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestArrayStack_New(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "create empty stack",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack()
			got := stack.List()
			if len(got) != len(tt.want) {
				t.Errorf("NewArrayStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Push(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "push to empty stack",
			data: []int{1},
			want: []int{1},
		},
		{
			name: "push multiple items",
			data: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack()
			for _, v := range tt.data {
				stack.Push(v)
			}
			got := stack.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After Push() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestArrayStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		init []int
		pops int
		want []int
	}{
		{
			name: "pop from empty stack",
			init: []int{},
			pops: 1,
			want: []int{},
		},
		{
			name: "pop single item",
			init: []int{1},
			pops: 1,
			want: []int{},
		},
		{
			name: "pop multiple items",
			init: []int{1, 2, 3},
			pops: 2,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack()
			for _, v := range tt.init {
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
}

func TestArrayStack_Print(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *ArrayStack
		expected string
	}{
		{
			name: "print_empty_stack",
			setup: func() *ArrayStack {
				return NewArrayStack()
			},
			expected: "print : \n",
		},
		{
			name: "print_single_element",
			setup: func() *ArrayStack {
				stack := NewArrayStack()
				stack.Push(1)
				return stack
			},
			expected: "print : 1 \n",
		},
		{
			name: "print_multiple_elements",
			setup: func() *ArrayStack {
				stack := NewArrayStack()
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
				return stack
			},
			expected: "print : 1 2 3 \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.setup()
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			stack.Print()

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			result := buf.String()

			if result != tt.expected {
				t.Errorf("Print() got %q, want %q", result, tt.expected)
			}
		})
	}
}

func BenchmarkArrayStack_Push(b *testing.B) {
	stack := NewArrayStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkArrayStack_Pop(b *testing.B) {
	stack := NewArrayStack()
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
