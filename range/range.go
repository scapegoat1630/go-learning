package _range

import (
	"math"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	if len(intervals) == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			if intervals[i][1] > ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	ans := make([][]int, 0)
	i := 0
	j := 0
	for i < len(A) && j < len(B) {
		l := A[i][0]
		r := A[i][1]
		if B[j][0] >= l {
			l = B[j][0]
		}
		if B[j][1] <= r {
			r = B[j][1]
		}
		if l <= r {
			ans = append(ans, []int{l, r})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

/*
给定一个 24 小时制（小时: 分钟"HH:MM"）的时间列表，找出列表中任意两个时间的
最小时间差并以分钟数表示。

1
输入: timePoints = ["23:59","00:00"]
输出: 1

2
输入: timePoints = ["00:00","23:59","00:00"]
输出 0
 */

func getMinutes(s string) int {
	return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'0')*10 + int(s[4]-'0')
}

func minTimeDiff(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	pre := getMinutes(timePoints[0])
	for i := 1; i < len(timePoints); i++ {
		diff := getMinutes(timePoints[i]) - pre
		if diff < ans {
			ans = diff
		}
		pre = getMinutes(timePoints[i])
	}
	diff := 24*60 - pre + getMinutes(timePoints[0])
	if ans > diff {
		ans = diff
	}
	return ans
}