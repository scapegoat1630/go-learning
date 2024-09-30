package search

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMin(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want int
		err  error
	}{
		{
			name: "test1",
			root: &Node{
				Val: 4,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 1,
					},
					Right: &Node{
						Val: 3,
					},
				},
				Right: &Node{
					Val: 6,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMin(tt.root)
			if got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimBTS(t *testing.T) {
	type args struct {
		n     *Node
		low   int
		high  int
		want1 *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				n: &Node{
					Val: 4,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 1,
						},
						Right: &Node{
							Val: 3,
						},
					},
					Right: &Node{
						Val: 6,
					},
				},
				low:   2,
				high:  4,
				want1: &Node{Val: 4, Left: &Node{Val: 2, Right: &Node{Val: 3}}},
			},
		},
		{
			name: "test2",
			args: args{
				n: &Node{
					Val: 4,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 1,
						},
						Right: &Node{
							Val: 3,
						},
					},
					Right: &Node{
						Val: 6,
					},
				},
				low:   5,
				high:  6,
				want1: &Node{Val: 6},
			},
		},
		{
			name: "test4",
			args: args{
				n:     nil,
				low:   1,
				high:  3,
				want1: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1 := trimBTS(tt.args.n, tt.args.low, tt.args.high); !reflect.DeepEqual(got1, tt.args.want1) {
				t.Errorf("trimBTS() = %v, want %v", got1, tt.args.want1)
			}
		})
	}
}

func Test_averageOfLevels(t *testing.T) {
	t.Parallel()
	type args struct {
		n *Node
	}
	tests := []struct {
		name         string
		args         args
		wantAverages []float64
	}{
		{
			name: "test1",
			args: args{
				n: &Node{
					Val: 4,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 1,
						},
						Right: &Node{
							Val: 3,
						},
					},
					Right: &Node{
						Val: 6,
					},
				},
			},
			wantAverages: []float64{4, 4, 2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotAverages := averageOfLevels(tt.args.n)
			fmt.Println(gotAverages)
			if !floatSliceEqual(gotAverages, tt.wantAverages) {
				t.Errorf("averageOfLevels() = %v, want %v", gotAverages, tt.wantAverages)
			}
		})
	}
}

func floatSliceEqual(f1, f2 []float64) bool {
	if len(f1) != len(f2) {
		return false
	}
	for i := range f1 {
		if f1[i] != f2[i] {
			return false
		}
	}
	return true
}
