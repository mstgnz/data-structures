package linkedlist

import (
	"testing"
)

func TestLinear_New(t *testing.T) {
	tests := []struct {
		name string
		data int
		want []int
	}{
		{
			name: "create with positive number",
			data: 1,
			want: []int{1},
		},
		{
			name: "create with zero",
			data: 0,
			want: []int{0},
		},
		{
			name: "create with negative number",
			data: -1,
			want: []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linear := NewLinear(tt.data)
			got := linear.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Linear() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinear_AddToAfter(t *testing.T) {
	tests := []struct {
		name    string
		init    []int
		data    int
		after   int
		want    []int
		wantErr bool
	}{
		{
			name:    "add after existing value",
			init:    []int{1, 2, 3},
			data:    4,
			after:   2,
			want:    []int{1, 2, 4, 3},
			wantErr: false,
		},
		{
			name:    "add after non-existing value",
			init:    []int{1, 2, 3},
			data:    4,
			after:   5,
			want:    []int{1, 2, 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linear := NewLinear(tt.init[0])
			for i := 1; i < len(tt.init); i++ {
				linear.AddToEnd(tt.init[i])
			}

			err := linear.AddToAfter(tt.data, tt.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddToAfter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := linear.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After AddToAfter() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinear_Delete(t *testing.T) {
	tests := []struct {
		name    string
		init    []int
		delete  int
		want    []int
		wantErr bool
	}{
		{
			name:    "delete existing value",
			init:    []int{1, 2, 3},
			delete:  2,
			want:    []int{1, 3},
			wantErr: false,
		},
		{
			name:    "delete non-existing value",
			init:    []int{1, 2, 3},
			delete:  4,
			want:    []int{1, 2, 3},
			wantErr: true,
		},
		{
			name:    "delete from single element list",
			init:    []int{1},
			delete:  1,
			want:    []int{0},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linear := NewLinear(tt.init[0])
			for i := 1; i < len(tt.init); i++ {
				linear.AddToEnd(tt.init[i])
			}

			err := linear.Delete(tt.delete)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := linear.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After Delete() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestLinear_Search(t *testing.T) {
	tests := []struct {
		name   string
		init   []int
		search int
		want   bool
	}{
		{
			name:   "search existing value",
			init:   []int{1, 2, 3},
			search: 2,
			want:   true,
		},
		{
			name:   "search non-existing value",
			init:   []int{1, 2, 3},
			search: 4,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linear := NewLinear(tt.init[0])
			for i := 1; i < len(tt.init); i++ {
				linear.AddToEnd(tt.init[i])
			}

			if got := linear.Search(tt.search); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinear_Print(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "print empty list",
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
			var list *Linear
			if len(tt.data) > 0 {
				list = NewLinear(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					list.AddToEnd(tt.data[i])
				}
			} else {
				list = NewLinear(0)
			}
			// Call Print method
			list.Print()
		})
	}
}

func BenchmarkLinear_AddToStart(b *testing.B) {
	linear := NewLinear(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linear.AddToStart(i)
	}
}

func BenchmarkLinear_AddToEnd(b *testing.B) {
	linear := NewLinear(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linear.AddToEnd(i)
	}
}

func BenchmarkLinear_Delete(b *testing.B) {
	linear := NewLinear(1)
	for i := 0; i < 1000; i++ {
		linear.AddToEnd(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = linear.Delete(i % 1000)
	}
}

func BenchmarkLinear_Search(b *testing.B) {
	linear := NewLinear(1)
	for i := 0; i < 1000; i++ {
		linear.AddToEnd(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linear.Search(i % 1000)
	}
}

func TestLinear_AddToStart(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      int
		expected []int
	}{
		{
			name:     "add to empty list",
			init:     0,
			add:      1,
			expected: []int{1, 0},
		},
		{
			name:     "add to non-empty list",
			init:     1,
			add:      2,
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinear(tt.init)
			list.AddToStart(tt.add)
			got := list.List()
			if len(got) != len(tt.expected) {
				t.Errorf("AddToStart() got %v, want %v", got, tt.expected)
				return
			}
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("AddToStart() got %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestLinear_AddToSequentially(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      []int
		expected []int
	}{
		{
			name:     "add sequentially to empty list",
			init:     0,
			add:      []int{1},
			expected: []int{0, 1},
		},
		{
			name:     "add sequentially multiple items",
			init:     1,
			add:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "add sequentially with duplicates",
			init:     1,
			add:      []int{2, 2, 3},
			expected: []int{1, 2, 2, 3},
		},
		{
			name:     "add sequentially with negative numbers",
			init:     0,
			add:      []int{-2, -1, 1},
			expected: []int{-2, -1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinear(tt.init)
			for _, v := range tt.add {
				list.AddToSequentially(v)
			}
			got := list.List()
			if len(got) != len(tt.expected) {
				t.Errorf("AddToSequentially() got %v, want %v", got, tt.expected)
				return
			}
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("AddToSequentially() got %v, want %v", got, tt.expected)
				}
			}
		})
	}
}
