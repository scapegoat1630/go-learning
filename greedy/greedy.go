package greedy

import "sort"

// Greedy algorithm

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	ans := 0
	add := false
	for _, v := range m {
		ans += v / 2 * 2
		if !add && v&1 == 1 {
			ans++
			add = true
		}
	}
	return ans
}

func leastInterval(tasks []byte, n int) int {
	hash := make(map[byte]int)
	for _, v := range tasks {
		hash[v]++
	}
	t := make([]int, 0)
	for _, v := range hash {
		t = append(t, v)
	}
	sort.Ints(t)
	ans := (n+1)*(t[len(t)-1]-1) + 1
	for i := len(t) - 2; i >= 0 && t[i] == t[len(t)-1]; i-- {
		ans++
	}
	return max(ans, len(tasks))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	n := len(nums)
	maxIdx := 0
	for i, j := range nums {
		if i < maxIdx {
			return false
		}

		maxIdx = max(maxIdx, i+j)
		if maxIdx >= n-1 {
			return true
		}

	}
	return false
}
