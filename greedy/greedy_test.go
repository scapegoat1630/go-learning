package greedy

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: "babad"},
			want: 5,
		},
		{
			name: "test2",
			args: args{s: "cbbd"},
			want: 3,
		},
		{
			name: "test3",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "test4",
			args: args{s: "ac"},
			want: 1,
		},
		{
			name: "test5",
			args: args{s: "abccccdd"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2},
			want: 8,
		},
		{
			name: "test2",
			args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 3},
			want: 9,
		},
		{
			name: "test3",
			args: args{tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'A', 'A', 'A', 'A', 'A', 'A'}, n: 5},
			want: 12,
		},
		{
			name: "test4",
			args: args{tasks: []byte{'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C'}, n: 3},
			want: 10,
		},
		{
			name: "test5",
			args: args{tasks: []byte{'A', 'B', 'A', 'B', 'A', 'B'}, n: 1},
			want: 6,
		},
		{
			name: "test6",
			args: args{tasks: []byte{'A', 'A', 'A', 'A', 'A'}, n: 0},
			want: 5,
		},
		{
			name: "test7",
			args: args{tasks: []byte{'A', 'A', 'A', 'A', 'A'}, n: 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "test2",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "test3",
			args: args{nums: []int{0}},
			want: true,
		},
		{
			name: "test4",
			args: args{nums: []int{1}},
			want: true,
		},
		{
			name: "test5",
			args: args{nums: []int{2, 0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
