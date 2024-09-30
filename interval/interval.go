package interval

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0)
	var curr []int
	for _, interval := range intervals {
		if len(curr) == 0 {
			curr = interval
			continue
		}
		if curr[1]+1 < interval[0] {
			res = append(res, curr)
			curr = interval
			continue
		}

		curr = []int{curr[0], max(interval[1], curr[1])}

	}
	if curr != nil {
		res = append(res, curr)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
