package array

import (
	"reflect"
	"testing"
)

func TestMergeArray(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeArray(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("MergeArray() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotateArray(t *testing.T) {
	type args struct {
		nums []int
		c    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				c:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				c:    0,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				c:    1,
			},
			want: []int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				c:    7,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "test5",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				c:    4,
			},
			want: []int{4, 5, 6, 7, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateArray(tt.args.nums, tt.args.c)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("RotateArray() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func TestArrayEqual(t *testing.T) {
	t.Log([26]int{1, 2, 3} == [26]int{1, 2, 3})
	c := 3
	a := [26]int{1, 2, 3}
	b := [26]int{1, 2, c}
	t.Log(a == b)
}
