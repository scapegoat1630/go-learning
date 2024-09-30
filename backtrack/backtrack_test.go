package backtrack

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "test1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			wantAns: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "test2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			wantAns: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "test4",
			args: args{
				candidates: []int{1},
				target:     1,
			},
			wantAns: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := combinationSum(tt.args.candidates, tt.args.target)
			sort.Ints(tt.wantAns[0])
			sort.Ints(gotAns[0])
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "test2",
			args: args{
				s: "0000",
			},
			want: []string{"0.0.0.0"},
		},
		{
			name: "test3",
			args: args{
				s: "101023",
			},
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreIpAddresses(tt.args.s)
			sort.Strings(tt.want)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countArrangement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				n: 7,
			},
			want: 34,
		},
		{
			name: "test3",
			args: args{
				n: 15,
			},
			want: 186,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangement(tt.args.n); got != tt.want {
				t.Errorf("countArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
