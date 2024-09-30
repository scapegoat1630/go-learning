package backtrack

import (
	"strconv"
	"strings"
)

// 回溯

func combinationSum(candidates []int, target int) (ans [][]int) {
	t := []int{}
	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, t...))
			return
		}
		if idx >= len(candidates) {
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			t = append(t, candidates[idx])
			dfs(target-candidates[idx], idx)
			t = t[:len(t)-1]
		}
	}
	dfs(target, 0)
	return
}

// 93.复原IP地址
func restoreIpAddresses(s string) []string {
	// 参数校验
	t := []string{}
	ans := []string{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			return
		}
		if len(t) == 3 {
			h := s[idx:]
			if helper(h) {
				ans = append(ans, strings.Join(append([]string{t[0], t[1], t[2]}, h), "."))
				return
			}

		}
		for i := 1; i <= 3 && idx+i < len(s); i++ {
			if helper(s[idx : idx+i]) {
				t = append(t, s[idx:idx+i])
				dfs(idx + i)
				t = t[:len(t)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func helper(s string) bool {
	// 最多3位
	if len(s) > 3 {
		return false
	}
	// 0x开头不行
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	//  <=255
	if num, _ := strconv.Atoi(s); num > 255 {
		return false
	}
	return true
}

// 526
func countArrangement(n int) int {
	ans := 0
	r := make([]bool, n+1)
	match := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}
	var dfs func(int)
	dfs = func(index int) {
		if index > n {
			ans++
			return
		}
		for _, v := range match[index] {
			if !r[v] {
				r[v] = true
				dfs(index + 1)
				r[v] = false
			}
		}
	}
	dfs(1)
	return ans
}
