package graph

import (
	"reflect"
	"testing"
)

func Test_watchedVideosByFriends(t *testing.T) {
	type args struct {
		watchedVideos [][]string
		friends       [][]int
		id            int
		level         int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-1",
			args: args{
				watchedVideos: [][]string{
					{"A", "B"},
					{"C"},
					{"B", "C"},
					{"D"},
				},
				friends: [][]int{
					{1, 2},
					{0, 3},
					{0, 3},
					{1, 2},
				},
				id:    0,
				level: 1,
			},
			want: []string{"B", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := watchedVideosByFriends(tt.args.watchedVideos, tt.args.friends, tt.args.id, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("watchedVideosByFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				matrix: [][]int{
					{9, 9, 4},
					{6, 6, 8},
					{2, 1, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
