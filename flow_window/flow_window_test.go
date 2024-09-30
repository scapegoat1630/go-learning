package flow_window

import (
	"reflect"
	"testing"
)

func TestMaxSubString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test1", "我是中国人", 5},
		{"test2", "你好", 2},
		{"test3", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubString(tt.s); got != tt.want {
				t.Errorf("MaxSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubString_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test1", "我是中国人", 5},
		{"test2", "bbbbb", 1},
		{"test3", "pwwkew", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubString(tt.s); got != tt.want {
				t.Errorf("MaxSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSubString(t *testing.T) {
	tests := []struct {
		name  string
		big   []int
		small []int
		want  []int
	}{
		{"test1", []int{1, 2, 3, 2, 4, 5, 1, 2, 3}, []int{1, 2, 3}, []int{0, 2}},
		{"test2", []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, []int{1, 5, 9}, []int{7, 10}},
		{"test3", []int{1, 2, 3}, []int{4}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubString(tt.big, tt.small); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
