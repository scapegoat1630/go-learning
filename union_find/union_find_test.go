package union_find

import (
	"reflect"
	"testing"
)

func TestGetLongest(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "get longest of single element",
			nums:   []int{1},
			expect: 1,
		},
		{
			name:   "get longest of ascending order",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 5,
		},
		{
			name:   "get longest of ascending order with gaps",
			nums:   []int{1, 3, 2, 4, 5},
			expect: 5,
		},
		{
			name:   "get longest of descending order",
			nums:   []int{5, 4, 3, 2, 1},
			expect: 5,
		},
		{
			name:   "get longest of descending order with gaps",
			nums:   []int{5, 6, 3, 2, 1, 4},
			expect: 6,
		},
		{
			name:   "get longest of random order",
			nums:   []int{5, 6, 3, 2, 1, 4, 8, 9, 10},
			expect: 6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetLongest(tc.nums)
			if got != tc.expect {
				t.Errorf("Expected: %d, got: %d", tc.expect, got)
			}
		})
	}
}

func TestGetRedundant(t *testing.T) {
	testCases := []struct {
		name   string
		edges  [][]int
		expect []int
	}{
		{
			name:   "get redundant edges of graph 1",
			edges:  [][]int{{1, 2}, {1, 3}, {2, 3}},
			expect: []int{2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetRedundant(tc.edges)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Expected: %v, got: %v", tc.expect, got)
			}
		})
	}
}
