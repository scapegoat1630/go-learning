package string

import (
	"strings"
)

func LongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	for i < len(strs[0]) {
		s := strs[0][i]
		for _, c := range strs[1:] {
			if c[i] != s || len(c)-1 < i {
				if i > 0 {
					return strs[0][0:i]
				} else {
					return ""
				}

			}
		}
		i++
	}
	return strs[0]
}

func rotateString(s, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

/*
*
给定两个字符串 s 和 p ，找到 s 中所有 p 的 异位词的子串，返回这些子串的起始索
引。不考虑答案输出的顺序。
异位词指由相同字母重排列形成的字符串（包括相同的字符串）。

1
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是"cba", 它是"abc" 的异位词。
起始索引等于 6 的子串是"bac", 它是"abc" 的异位词。

2
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是"ab", 它是"ab" 的异位词。
起始索引等于 1 的子串是"ba", 它是"ab" 的异位词。
起始索引等于 2 的子串是"ab", 它是"ab" 的异位词。
*/
func FindE(s, p string) (ans []int) {
	ans = []int{}
	if len(p) > len(s) {
		return ans
	}
	st := [26]int{}
	pt := [26]int{}
	for _, c := range p {
		pt[c-'a']++
	}
	for i, c := range s {
		st[c-'a']++
		if i >= len(p) {
			sc := s[i-len(p)] - 'a'
			st[sc]--
		}
		if st == pt {
			ans = append(ans, i-len(p)+1)
		}
	}
	return
}

func TimeRequireToBuy(tickets []int, k int) int {
	res := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			res += Min(tickets[k], tickets[i])
		} else {
			res += Min(tickets[k]-1, tickets[i])
		}
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
