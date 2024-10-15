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

// 6. Z 字形变换
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
//
// 示例 1：
//
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// 示例 3：
//
// 输入：s = "A", numRows = 1
// 输出："A"
//
// 提示：
//
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans := make([]string, numRows)
	for i, j := range s {
		idx := i % (2*numRows - 2)
		if idx >= numRows {
			idx = 2*numRows - idx - 2
		}
		ans[idx] = ans[idx] + string(j)
	}
	a := ""
	for _, j := range ans {
		a = a + j
	}
	return a
}
