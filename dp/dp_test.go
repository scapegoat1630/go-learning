package dp

import (
	"reflect"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_climbStairs_1",
			args: args{2},
			want: 2,
		},
		{
			name: "test_climbStairs_2",
			args: args{5},
			want: 8,
		},
		{
			name: "test_climbStairs_3",
			args: args{0},
			want: 1,
		},
		{
			name: "test_climbStairs_4",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_rob_1",
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "test_rob_2",
			args: args{[]int{2, 7, 9, 3, 1}},
			want: 24,
		},
		{
			name: "test_rob_3",
			args: args{[]int{2, 1, 1, 2}},
			want: 4,
		},
		{
			name: "test_rob_4",
			args: args{[]int{1, 2, 3}},
			want: 4,
		},
		{
			name: "test_rob_5",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "test_rob_6",
			args: args{[]int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestSufficientTeam(t *testing.T) {
	type args struct {
		skills []string
		people [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_smallestSufficientTeam_1",
			args: args{[]string{"java", "go", "react"}, [][]string{{"java"}, {"go"}, {"react"}, {"js"}}},
			want: []int{0, 1, 2},
		},
		{
			name: "test_smallestSufficientTeam_2",
			args: args{[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws", "linux"}, [][]string{{"algorithms", "math", "java"}, {"java", "csharp", "reactjs"}, {"csharp", "math", "linux"}, {"linux", "reactjs"}}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "test_smallestSufficientTeam_3",
			args: args{[]string{"a", "b", "c", "d"}, [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"a", "c", "d"}, {"a", "b", "d"}}},
			want: []int{0, 1, 3},
		},
		{
			name: "test_smallestSufficientTeam_4",
			args: args{[]string{"a", "b", "c", "d"}, [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"a", "c", "d"}, {"a", "b", "d"}, {"a", "b", "c", "d"}}},
			want: []int{0},
		},
		{
			name: "test_smallestSufficientTeam_5",
			args: args{[]string{"a", "b", "c", "d"}, [][]string{{"a", "b", "c", "d"}}},
			want: []int{0},
		},
		{
			name: "test_smallestSufficientTeam_6",
			args: args{[]string{"a", "b", "c", "d"}, [][]string{}},
			want: []int{},
		},
		{
			name: "test_smallestSufficientTeam_7",
			args: args{[]string{}, [][]string{{"a", "b", "c", "d"}}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSufficientTeam(tt.args.skills, tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSufficientTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
