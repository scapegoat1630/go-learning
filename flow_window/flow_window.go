package flow_window

import "math"

func MaxSubString(s string) int {
	t := []rune(s)
	l, r := 0, 0
	m := make(map[rune]int)
	ans := 0
	for r < len(t) {
		idx, ok := m[t[r]]
		if !ok {
			if r-l+1 > ans {
				ans = r - l + 1
			}
		} else {
			delete(m, t[r])
			l = idx + 1
		}
		m[t[r]] = r
		r++
	}
	return ans
}

func MinSubString(big []int, small []int) []int {
	smallMap := make(map[int]int, 0)
	windowMap := make(map[int]int, 0)
	smallCount := len(small)
	minLen := math.MaxInt
	count := 0
	for _, v := range small {
		smallMap[v]++
	}
	left, right := 0, 0
	resultLeft, resultRight := 0, 0
	for right < len(big) {
		w, ok := smallMap[big[right]]
		if !ok {
			right++
			continue
		}
		windowMap[big[right]]++
		if windowMap[big[right]] <= w {
			count++
		}
		for count == smallCount {
			if right-left+1 < minLen {
				minLen = right - left + 1
				resultLeft, resultRight = left, right
			}
			_, ok := windowMap[big[left]]
			if ok {
				windowMap[big[left]]--
				if windowMap[big[left]] < smallMap[big[left]] {
					count--
				}
			}
			left++
		}
		right++
	}
	if minLen == math.MaxInt {
		return []int{}
	}
	return []int{resultLeft, resultRight}

}
