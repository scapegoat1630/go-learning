package _range

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				intervals: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "test2",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test3",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				A: [][]int{{1, 3}, {5, 9}},
				B: [][]int{{2, 4}, {6, 8}},
			},
			want: [][]int{{2, 3}, {6, 8}},
		},
		{
			name: "test2",
			args: args{
				A: [][]int{{1, 5}},
				B: [][]int{{2, 3}, {4, 7}},
			},
			want: [][]int{{2, 3}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minTimeDiff(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				timePoints: []string{"23:59", "00:00"},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				timePoints: []string{"23:59", "00:00", "12:00"},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				timePoints: []string{"11:00", "12:00", "13:00"},
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeDiff(tt.args.timePoints); got != tt.want {
				t.Errorf("minTimeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
