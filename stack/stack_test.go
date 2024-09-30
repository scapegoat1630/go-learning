package stack

import (
	"reflect"
	"testing"
)

func TestValidBracket(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "test1",
			s:    "{[()]}",
			want: true,
		},
		{
			name: "test2",
			s:    "{[(])}",
			want: false,
		},
		{
			name: "test3",
			s:    "{{[[(())]]}}",
			want: true,
		},
		{
			name: "test4",
			s:    "{{[[(())]]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidBracket(tt.s); got != tt.want {
				t.Errorf("ValidBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		want   int
		isFail bool
	}{
		{
			name:   "test1",
			nums:   []int{7, 1, 5, 3, 6, 4},
			want:   5,
			isFail: false,
		},
		{
			name:   "test2",
			nums:   []int{7, 6, 4, 3, 1},
			want:   0,
			isFail: false,
		},
		{
			name:   "test3",
			nums:   []int{2, 1, 3, 4, 5, 6, 7},
			want:   6,
			isFail: false,
		},
		{
			name:   "test4",
			nums:   []int{1, 2, 3, 4, 5},
			want:   4,
			isFail: false,
		},
		{
			name:   "test5",
			nums:   []int{1, 2},
			want:   1,
			isFail: false,
		},
		{
			name:   "test6",
			nums:   []int{},
			want:   0,
			isFail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.nums)
			if got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextBigNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			name: "test2",
			args: args{
				nums1: []int{4, 3, 1},
				nums2: []int{9, 2, 3, 4, 5, 6, 7, 8, 1},
			},
			want: []int{5, 4, -1},
		},
		{
			name: "test3",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextBigNumber(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextBigNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargeArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "test1",
			height: []int{2, 1, 2},
			want:   3,
		},
		{
			name:   "test2",
			height: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "test3",
			height: []int{2, 1, 5, 6, 2, 3},
			want:   10,
		},
		{
			name:   "test4",
			height: []int{2, 4},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.height)
			if got != tt.want {
				t.Errorf("LargeArea() = %v, want %v", got, tt.want)
			}
			got = LargeArea(tt.height)
			if got != tt.want {
				t.Errorf("LargeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
