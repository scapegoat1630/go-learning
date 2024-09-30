package tree

import (
	"reflect"
	"testing"
)

func TestInOrder(t *testing.T) {
	tests := []struct {
		name  string
		input *Node
		want  []int
	}{
		{
			"example 1",
			&Node{1, nil, nil},
			[]int{1},
		},
		{
			"example 2",
			&Node{1, &Node{2, nil, nil}, &Node{3, nil, nil}},
			[]int{2, 1, 3},
		},
		{
			"example 3",
			&Node{1, &Node{2, nil, nil}, &Node{3, &Node{4, nil, nil}, nil}},
			[]int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrder(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			"example 1",
			0,
			1,
		},
		{
			"example 2",
			1,
			1,
		},
		{
			"example 3",
			2,
			2,
		},
		{
			"example 4",
			3,
			5,
		},
		{
			"example 5",
			4,
			14,
		},
		{
			"example 6",
			5,
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees(tt.n); got != tt.want {
				t.Errorf("NumTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name  string
		input *Node
		want  []int
	}{
		{
			"example 1",
			&Node{
				Val: 1,
				Right: &Node{
					Val: 2,
					Right: &Node{
						Val: 5,
					},
					Left: &Node{
						Val:   3,
						Right: &Node{Val: 4},
					},
				},
			},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"example 2",
			&Node{
				Val: 1,
			},
			[]int{1},
		},
		{
			"example 3",
			&Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val: 3,
				},
			},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToBinarySearchTree(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		want    *Node
		wantErr bool
	}{
		{
			"example 1",
			[]int{1},
			&Node{1, nil, nil},
			false,
		},
		{
			"example 2",
			[]int{1, 2, 3},
			&Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}},
			false,
		},
		{
			"example 3",
			[]int{1, 2, 3, 4},
			&Node{2, &Node{1, nil, nil}, &Node{4, nil, &Node{3, nil, nil}}},
			false,
		},
		{
			"error case: empty array",
			[]int{},
			nil,
			false,
		},
		{
			"error case: single element",
			[]int{1},
			&Node{1, nil, nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayToBinarySearchTree(tt.nums)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToBinarySearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
