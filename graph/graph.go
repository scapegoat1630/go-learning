package graph

import (
	"fmt"
	"sort"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	// 找邻居
	q := []int{id}
	m := make(map[int]bool)
	m[id] = true
	for i := 0; i < level; i++ {
		newQ := []int{}
		for _, v := range q {
			for _, n := range friends[v] {
				if !m[n] {
					newQ = append(newQ, n)
					m[n] = true
				}
			}
		}
		q = newQ
	}
	fmt.Println(q)
	// 找电影
	v := make(map[string]int)
	for _, t := range q {
		for _, w := range watchedVideos[t] {
			v[w]++
		}
	}
	fmt.Println(v)
	// 排序
	s := make([]string, 0)
	for k, _ := range v {
		s = append(s, k)
	}
	sort.Slice(s, func(i, j int) bool {
		return v[s[i]] < v[s[j]] || (v[s[i]] == v[s[j]] && s[i] < s[j])
	})
	return s
}

func findCenter(edges [][]int) int {
	m := make(map[int]int)
	for _, v := range edges {
		m[v[0]]++
		m[v[1]]++
	}
	for k, v := range m {
		if v == len(edges)-1 {
			return k
		}
	}
	return -1
}

func findCenter2(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				ans = max(ans, dfs(matrix, dp, i, j))
			}
		}
	}
	return ans
}
func dfs(matrix [][]int, dp [][]int, x, y int) int {
	m, n := len(matrix), len(matrix[0])
	if dp[x][y] != 0 {
		return dp[x][y]
	}
	dp[x][y] = 1
	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx, ny := x+d[0], y+d[1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		if matrix[nx][ny] > matrix[x][y] {
			dp[x][y] = max(dfs(matrix, dp, nx, ny)+1, dp[x][y])
		}
	}
	return dp[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
