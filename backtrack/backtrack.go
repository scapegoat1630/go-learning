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

// 46. 全排列
// https://leetcode.cn/problems/permutations/description/
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
//
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同

func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	tmp := append([]int{}, nums...)
	var dfs func(int)
	dfs = func(first int) {
		if first == n {
			t := append([]int{}, tmp...)
			ans = append(ans, t)
			return
		}
		for i := first; i < n; i++ {
			tmp[i], tmp[first] = tmp[first], tmp[i]
			dfs(first + 1)
			tmp[i], tmp[first] = tmp[first], tmp[i]
		}
	}
	dfs(0)
	return ans
}
