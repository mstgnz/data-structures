package linkedlist

import (
	"reflect"
	"testing"
)

func TestCircular_New(t *testing.T) {
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
			list := Circular(tt.data)
			got := list.List()
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Circular() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCircular_AddToSequentially(t *testing.T) {
	tests := []struct {
		name string
		init int
		add  []int
		want []int
	}{
		{
			name: "add sequentially to empty list",
			init: 0,
			add:  []int{1},
			want: []int{0, 1},
		},
		{
			name: "add sequentially multiple items",
			init: 1,
			add:  []int{2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "add sequentially with duplicates",
			init: 1,
			add:  []int{2, 2, 3},
			want: []int{1, 2, 2, 3},
		},
		{
			name: "add sequentially with negative numbers",
			init: 0,
			add:  []int{-2, -1, 1},
			want: []int{-2, -1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Circular(tt.init)
			for _, v := range tt.add {
				list.AddToSequentially(v)
			}
			got := list.List()
			if len(got) != len(tt.want) {
				t.Errorf("After AddToSequentially() = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After AddToSequentially() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCircular_Delete(t *testing.T) {
	tests := []struct {
		name    string
		init    int
		setup   []int
		delete  int
		want    []int
		wantErr bool
	}{
		{
			name:    "delete from single element list",
			init:    1,
			setup:   []int{},
			delete:  1,
			want:    []int{0},
			wantErr: false,
		},
		{
			name:    "delete first element",
			init:    1,
			setup:   []int{2, 3},
			delete:  1,
			want:    []int{2, 3},
			wantErr: false,
		},
		{
			name:    "delete middle element",
			init:    1,
			setup:   []int{2, 3},
			delete:  2,
			want:    []int{1, 3},
			wantErr: false,
		},
		{
			name:    "delete last element",
			init:    1,
			setup:   []int{2, 3},
			delete:  3,
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name:    "delete non-existing element",
			init:    1,
			setup:   []int{2, 3},
			delete:  4,
			want:    []int{1, 2, 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Circular(tt.init)
			for _, v := range tt.setup {
				list.AddToEnd(v)
			}
			err := list.Delete(tt.delete)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := list.List()
			if len(got) != len(tt.want) {
				t.Errorf("After Delete() = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("After Delete() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCircular_Print(t *testing.T) {
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
			data: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list ICircular
			if len(tt.data) > 0 {
				list = Circular(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					list.AddToEnd(tt.data[i])
				}
			} else {
				list = Circular(0)
			}
			list.Print()
		})
	}
}

func TestCircular_AddToEnd(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      []int
		expected []int
	}{
		{
			name:     "add to empty list",
			init:     0,
			add:      []int{1},
			expected: []int{0, 1},
		},
		{
			name:     "add multiple items",
			init:     1,
			add:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Circular(tt.init)
			for _, v := range tt.add {
				list.AddToEnd(v)
			}
			got := list.List()
			if len(got) != len(tt.expected) {
				t.Errorf("AddToEnd() got %v, want %v", got, tt.expected)
				return
			}
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("AddToEnd() got %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestCircular(t *testing.T) {
	circular := Circular(1)
	expect := []int{1}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Circular() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToAfter(t *testing.T) {
	circular := Circular(1)
	circular.AddToAfter(2, 1)
	expect := []int{1, 2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToAfter() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToEnd(t *testing.T) {
	circular := Circular(1)
	circular.AddToEnd(2)
	expect := []int{1, 2}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToEnd() = %v, want %v", got, expect)
	}
}

func Test_circular_AddToStart(t *testing.T) {
	circular := Circular(1)
	circular.AddToStart(2)
	expect := []int{2, 1}
	if got := circular.List(); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToStart() = %v, want %v", got, expect)
	}
}

func TestCircular_AddToAfter(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() ICircular
		data     int
		after    int
		expected []int
	}{
		{
			name: "add_after_empty_list",
			setup: func() ICircular {
				return Circular(0)
			},
			data:     2,
			after:    0,
			expected: []int{0, 2},
		},
		{
			name: "add_after_single_element",
			setup: func() ICircular {
				list := Circular(1)
				return list
			},
			data:     2,
			after:    1,
			expected: []int{1, 2},
		},
		{
			name: "add_after_multiple_elements",
			setup: func() ICircular {
				list := Circular(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			data:     4,
			after:    2,
			expected: []int{1, 2, 4, 3},
		},
		{
			name: "add_after_non_existent_element",
			setup: func() ICircular {
				list := Circular(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			data:     4,
			after:    5,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			list.AddToAfter(tt.data, tt.after)
			result := list.List()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddToAfter() got %v, want %v", result, tt.expected)
			}
		})
	}
}
