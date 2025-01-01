package LinkedList

import (
	"reflect"
	"testing"
)

func TestDouble_New(t *testing.T) {
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
			list := Double(tt.data)
			got := list.List(false)
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Double() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDouble_Delete(t *testing.T) {
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
			list := Double(tt.init)
			for _, v := range tt.setup {
				list.AddToEnd(v)
			}
			err := list.Delete(tt.delete)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := list.List(false)
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

func TestDouble_Print(t *testing.T) {
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
			var list IDouble
			if len(tt.data) > 0 {
				list = Double(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					list.AddToEnd(tt.data[i])
				}
			} else {
				list = Double(0)
			}
			list.Print(false)
			list.Print(true)
		})
	}
}

func TestDouble(t *testing.T) {
	double := Double(1)
	expect := []int{1}
	if got := double.List(true); !reflect.DeepEqual(got, expect) {
		t.Errorf("Double() = %v, want %v", got, expect)
	}
}

func Test_double_AddToAfter(t *testing.T) {
	double := Double(1)
	double.AddToAfter(2, 1)
	expect := []int{1, 2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToAfter() = %v, want %v", got, expect)
	}
}

func Test_double_AddToEnd(t *testing.T) {
	double := Double(1)
	double.AddToEnd(2)
	expect := []int{1, 2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToEnd() = %v, want %v", got, expect)
	}
}

func Test_double_AddToSequentially(t *testing.T) {
	double := Double(1)
	double.AddToSequentially(2)
	expect := []int{1, 2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToSequentially() = %v, want %v", got, expect)
	}
}

func Test_double_AddToStart(t *testing.T) {
	double := Double(1)
	double.AddToStart(2)
	expect := []int{2, 1}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("AddToStart() = %v, want %v", got, expect)
	}
}

func Test_double_Delete(t *testing.T) {
	double := Double(1)
	double.AddToStart(2)
	double.Delete(1)
	expect := []int{2}
	if got := double.List(false); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}

func TestDouble_AddToStart(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() IDouble
		data     int
		expected []int
	}{
		{
			name: "add_to_start_empty_list",
			setup: func() IDouble {
				return Double(0)
			},
			data:     1,
			expected: []int{1, 0},
		},
		{
			name: "add_to_start_single_element",
			setup: func() IDouble {
				list := Double(1)
				return list
			},
			data:     2,
			expected: []int{2, 1},
		},
		{
			name: "add_to_start_multiple_elements",
			setup: func() IDouble {
				list := Double(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			data:     4,
			expected: []int{4, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			list.AddToStart(tt.data)
			result := list.List(false)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddToStart() got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDouble_AddToSequentially(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() IDouble
		data     int
		expected []int
	}{
		{
			name: "add_sequentially_empty_list",
			setup: func() IDouble {
				return Double(0)
			},
			data:     1,
			expected: []int{0, 1},
		},
		{
			name: "add_sequentially_single_element",
			setup: func() IDouble {
				list := Double(1)
				return list
			},
			data:     2,
			expected: []int{1, 2},
		},
		{
			name: "add_sequentially_multiple_elements",
			setup: func() IDouble {
				list := Double(1)
				list.AddToEnd(3)
				list.AddToEnd(5)
				return list
			},
			data:     2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name: "add_sequentially_duplicate_value",
			setup: func() IDouble {
				list := Double(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			data:     2,
			expected: []int{1, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			list.AddToSequentially(tt.data)
			result := list.List(false)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddToSequentially() got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDouble_AddToAfter(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() IDouble
		after    int
		data     int
		expected []int
	}{
		{
			name: "add_after_empty_list",
			setup: func() IDouble {
				return Double(0)
			},
			after:    0,
			data:     1,
			expected: []int{0, 1},
		},
		{
			name: "add_after_single_element",
			setup: func() IDouble {
				list := Double(1)
				return list
			},
			after:    1,
			data:     2,
			expected: []int{1, 2},
		},
		{
			name: "add_after_multiple_elements",
			setup: func() IDouble {
				list := Double(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			after:    2,
			data:     4,
			expected: []int{1, 2, 4, 3},
		},
		{
			name: "add_after_non_existent_element",
			setup: func() IDouble {
				list := Double(1)
				list.AddToEnd(2)
				list.AddToEnd(3)
				return list
			},
			after:    4,
			data:     5,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			list.AddToAfter(tt.data, tt.after)
			result := list.List(false)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddToAfter() got %v, want %v", result, tt.expected)
			}
		})
	}
}
