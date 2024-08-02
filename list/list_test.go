package list

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input *Node
		want  *Node
	}{
		{
			name:  "reverse empty list",
			input: nil,
			want:  nil,
		},
		{
			name:  "reverse list with one node",
			input: &Node{Val: 1},
			want:  &Node{Val: 1},
		},
		{
			name:  "reverse list with two nodes",
			input: &Node{Val: 1, Next: &Node{Val: 2}},
			want:  &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			name:  "reverse list with three nodes",
			input: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
			want:  &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			for got != nil {
				if tt.want == nil {
					t.Error("got non-nil Node")
				}
				if got.Val != tt.want.Val {
					t.Errorf("got %d; want %d", got.Val, tt.want.Val)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
			if tt.want != nil {
				t.Errorf("got nil; want %d", tt.want.Val)
			}
		})
	}
}

func TestReverse2(t *testing.T) {
	tests := []struct {
		name  string
		input *Node
		want  *Node
	}{
		{
			name:  "reverse empty list",
			input: nil,
			want:  nil,
		},
		{
			name:  "reverse list with one node",
			input: &Node{Val: 1},
			want:  &Node{Val: 1},
		},
		{
			name:  "reverse list with two nodes",
			input: &Node{Val: 1, Next: &Node{Val: 2}},
			want:  &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			name:  "reverse list with three nodes",
			input: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
			want:  &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse2(tt.input)
			for got != nil {
				if tt.want == nil {
					t.Error("got non-nil Node")
				}
				if got.Val != tt.want.Val {
					t.Errorf("got %d; want %d", got.Val, tt.want.Val)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
			if tt.want != nil {
				t.Errorf("got nil; want %d", tt.want.Val)
			}
		})
	}
}
