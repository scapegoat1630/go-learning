package string

import (
	"reflect"
	"testing"
)

func TestLongestPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{}, ""},
		{[]string{"flower"}, "flower"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, test := range tests {
		got := LongestPrefix(test.strs)
		if got != test.want {
			t.Errorf("LongestPrefix(%v) = %v, want %v", test.strs, got, test.want)
		}
	}
}

func TestFindE(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		s       string
		p       string
		wantAns []int
	}{
		{
			"first test",
			"cbaebabacd",
			"abc",
			[]int{0, 6},
		},
		{
			"second test",
			"abab",
			"ab",
			[]int{0, 1, 2},
		},
		{
			"third test",
			"acb",
			"ab",
			[]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotAns := FindE(tt.s, tt.p)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("FindE() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_timeRequireToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				tickets: []int{1, 3, 2},
				k:       1,
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				tickets: []int{2, 2, 3, 3, 2, 2, 1, 1, 2, 3},
				k:       3,
			},
			want: 20,
		},
		{
			name: "test3",
			args: args{
				tickets: []int{1, 2, 3, 2, 5, 1, 3, 2, 1},
				k:       4,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeRequireToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("TimeRequireToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
