package interval

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
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
				intervals: [][]int{{1, 3}, {2, 6}, {3, 4}, {5, 7}},
			},
			want: [][]int{{1,  7}},
		},
		{
			name: "test2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "test3",
			args: args{
				intervals: [][]int{{1, 4}, {2, 6}},
			},
			want: [][]int{{1, 6}},
		},
		{
			name: "test4",
			args: args{
				intervals: [][]int{{1, 5}, {2, 3}, {3, 4}, {4, 6}},
			},
			want: [][]int{{1, 6}},
		},
		{
			name: "test5",
			args: args{
				intervals: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "test6",
			args: args{
				intervals: [][]int{{1, 3}},
			},
			want: [][]int{{1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
