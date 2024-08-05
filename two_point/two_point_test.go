package two_point

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		desc string
	}{
		{
			"test1",
			[]int{1, 0, 2, 0, 3},
			[]int{1, 2, 3, 0, 0},
			"Move all zeros to the end",
		},
		{
			"test2",
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
			"Move all zeros to the end",
		},
		{
			"test3",
			[]int{1, 2, 0, 0, 3},
			[]int{1, 2, 3, 0, 0},
			"Move all zeros to the end",
		},
		{
			"test4",
			[]int{0, 1, 2, 0, 3},
			[]int{1, 2, 3, 0, 0},
			"Move all zeros to the end",
		},
		{
			"test5",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			"No zeros in the array",
		},
		{
			"test6",
			[]int{1, 2, 3, 4, 0},
			[]int{1, 2, 3, 4, 0},
			"No zeros in the array",
		},
		{
			"test7",
			[]int{0, 1, 2, 3, 4},
			[]int{1, 2, 3, 4, 0},
			"Only zeros in the array",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeros(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("MoveZeros() = %v, want %v, test case: %s", tt.nums, tt.want, tt.desc)
			}
		})
	}
}

func Test_subSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"test1",
			[]int{1, 0, 2, 0, 3},
			[]int{0, 3},
		},
		{
			"test2",
			[]int{0, 1, 0, 3, 12},
			[]int{1, 2},
		},
		{
			"test3",
			[]int{1, 2, 0, 0, 3},
			[]int{0, 3},
		},
		{
			"test4",
			[]int{0, 1, 2, 0, 3},
			[]int{1, 3},
		},
		{
			"test5",
			[]int{1, 2, 3, 4, 5},
			[]int{-1, -1},
		},
		{
			"test6",
			[]int{1, 2, 3, 4, 0},
			[]int{0, 4},
		},
		{
			"test7",
			[]int{0, 1, 2, 3, 4},
			[]int{-1, -1},
		},
		{
			"test8",
			[]int{4, 3, 2, 1, 0},
			[]int{0, 4},
		},
		{
			"test9",
			[]int{5, 4, 3, 2, 1, 0},
			[]int{0, 5},
		},
		{
			"test10",
			[]int{0, 0, 0, 0, 0},
			[]int{-1, -1},
		},
		{
			"test11",
			[]int{1},
			[]int{-1, -1},
		},
		{
			"test12",
			[]int{-1, -2, -3, -4, -5},
			[]int{0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subSort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pairSums(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			"test1",
			[]int{1, 0, 2, 0, 3},
			3,
			[][]int{{1, 2}, {0, 3}},
		},
		{
			"test2",
			[]int{0, 1, 0, 3, 12},
			4,
			[][]int{{1, 3}},
		},
		{
			"test3",
			[]int{5, 6, 5},
			11,
			[][]int{{5, 6}},
		},
		{
			"test3",
			[]int{5, 6, 5, 6},
			11,
			[][]int{{5, 6}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSums(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pairSums2(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			"test1",
			[]int{1, 0, 2, 0, 3},
			3,
			[][]int{{0, 3}, {1, 2}},
		},
		{
			"test2",
			[]int{0, 1, 0, 3, 12},
			4,
			[][]int{{1, 3}},
		},
		{
			"test3",
			[]int{5, 6, 5},
			11,
			[][]int{{5, 6}},
		},
		{
			"test3",
			[]int{5, 6, 5, 6},
			11,
			[][]int{{5, 6}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum2(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"test1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.nums); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"test1",
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			"test2",
			[]int{4, 2, 3},
			1,
		},
		{
			"test3",
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.nums); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{[]int{2, 0, 2, 1, 1, 0}},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"test2",
			args{[]int{2, 0, 1}},
			[]int{0, 1, 2},
		},
		{
			"test3",
			args{[]int{0}},
			[]int{0},
		},
		{
			"test4",
			args{[]int{1}},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
