package list

import (
	"github.com/stretchr/testify/assert"
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

func TestReverse3(t *testing.T) {
	cases := []struct {
		name     string
		head     *Node
		left     int
		right    int
		expected *Node
	}{
		{
			name:     "reverse empty list",
			head:     nil,
			left:     1,
			right:    1,
			expected: nil,
		},
		{
			name:     "reverse list with one node",
			head:     &Node{Val: 1},
			left:     1,
			right:    1,
			expected: &Node{Val: 1},
		},
		{
			name:     "reverse list with two nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2}},
			left:     1,
			right:    2,
			expected: &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			name:     "reverse list with three nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
			left:     1,
			right:    3,
			expected: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}},
		},
		{
			name:     "reverse list with many nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
			left:     3,
			right:    5,
			expected: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 5, Next: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Reverse3(tc.head, tc.left, tc.right)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	cases := []struct {
		name     string
		head     *Node
		left     int
		right    int
		expected *Node
	}{
		{
			name:     "reverse empty list",
			head:     nil,
			left:     1,
			right:    1,
			expected: nil,
		},
		{
			name:     "reverse list with one node",
			head:     &Node{Val: 1},
			left:     1,
			right:    1,
			expected: &Node{Val: 1},
		},
		{
			name:     "reverse list with two nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2}},
			left:     1,
			right:    2,
			expected: &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			name:     "reverse list with three nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
			left:     1,
			right:    3,
			expected: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}},
		},
		{
			name:     "reverse list with many nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
			left:     3,
			right:    5,
			expected: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 5, Next: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverseBetween(tc.head, tc.left, tc.right)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_rotateRight(t *testing.T) {
	cases := []struct {
		name     string
		head     *Node
		k        int
		expected *Node
	}{
		{
			name:     "rotate empty list",
			head:     nil,
			k:        1,
			expected: nil,
		},
		{
			name:     "rotate list with one node",
			head:     &Node{Val: 1},
			k:        1,
			expected: &Node{Val: 1},
		},
		{
			name:     "rotate list with two nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2}},
			k:        0,
			expected: &Node{Val: 1, Next: &Node{Val: 2}},
		},
		{
			name:     "rotate list with many nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
			k:        4,
			expected: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}}}}}},
		},
		{
			name:     "rotate list with many nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
			k:        6,
			expected: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8, Next: &Node{Val: 1, Next: &Node{Val: 2}}}}}}}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rotateRight(tc.head, tc.k)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_partition(t *testing.T) {
	cases := []struct {
		name     string
		head     *Node
		x        int
		expected *Node
	}{
		{
			name:     "partition empty list",
			head:     nil,
			x:        1,
			expected: nil,
		},
		{
			name:     "partition list with one node",
			head:     &Node{Val: 1},
			x:        1,
			expected: &Node{Val: 1},
		},
		{
			name:     "partition list with two nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2}},
			x:        1,
			expected: &Node{Val: 1, Next: &Node{Val: 2}},
		},
		{
			name:     "partition list with many nodes",
			head:     &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
			x:        4,
			expected: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6, Next: &Node{Val: 7, Next: &Node{Val: 8}}}}}}}},
		},
		{
			name:     "partition list with many nodes",
			head:     &Node{Val: 8, Next: &Node{Val: 7, Next: &Node{Val: 6, Next: &Node{Val: 5, Next: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}}}}}}},
			x:        4,
			expected: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1, Next: &Node{Val: 8, Next: &Node{Val: 7, Next: &Node{Val: 6, Next: &Node{Val: 5}}}}}}}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partition(tc.head, tc.x)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_hasCycle(t *testing.T) {
	t.Parallel()
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has cycle",
			args: args{head: &Node{
				Val: 1,
				Next: &Node{
					Val:  2,
					Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 1}}}},
				},
			}},
			want: true,
		},
		{
			name: "no cycle",
			args: args{head: &Node{
				Val: 1,
				Next: &Node{
					Val:  2,
					Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}},
				},
			}},
			want: false,
		},
		{
			name: "nil node",
			args: args{head: nil},
			want: false,
		},
		{
			name: "single node",
			args: args{head: &Node{Val: 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	// 示例代码，创建一个链表并检查是否为回文
	// 链表为 1 -> 2 -> 2 -> 1
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 2}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	result := isPalindrome(node1)
	println(result)
}
