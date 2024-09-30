package stack

import (
	"fmt"
	"math"
)

func ValidBracket(s string) bool {
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]string, 0)
	for _, s := range s {
		p, ok := m[string(s)]
		if !ok {
			stack = append(stack, string(s))
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != p {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func MaxProfit(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	maxProfit := math.MinInt32
	for i := 1; i < len(nums); i++ {
		if min >= nums[i] {
			min = nums[i]
			continue
		}
		if nums[i]-min > maxProfit {
			maxProfit = nums[i] - min
		}
	}
	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}

func NextBigNumber(nums1, nums2 []int) (ans []int) {
	if len(nums2) == 0 || len(nums1) == 0 {
		return []int{}
	}
	if len(nums2) == 1 {
		return []int{-1}
	}
	var (
		m = make(map[int]int)
		s = make([]int, 0)
		t = make([]int, 0, len(nums2))
	)
	for _, _ = range nums2 {
		t = append(t, -1)
	}
	for k, v := range nums2 {
		m[v] = k
		for len(s) > 0 {
			if nums2[s[len(s)-1]] < v {
				t[s[len(s)-1]] = k
				s = s[:len(s)-1]
			} else {
				break
			}
		}
		s = append(s, k)
	}
	// fmt.Println(t)
	for _, v := range nums1 {
		if m[v] == -1 {
			ans = append(ans, -1)
		} else {
			a := t[m[v]]
			if a >= 0 {
				ans = append(ans, nums2[a])
			} else {
				ans = append(ans, -1)
			}

		}
	}
	return ans

}
func LargeArea(height []int) (area int) {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] <= height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	fmt.Println(left)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && height[i] <= height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	fmt.Println(right)
	for i := 0; i < n; i++ {
		area = Max(area, (right[i]-left[i]-1)*height[i])
	}
	return area
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	fmt.Println(left)
	mono_stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	fmt.Println(right)
	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans

}

func Max(nums1 int, nums2 int) int {
	if nums1 > nums2 {
		return nums1
	} else {
		return nums2
	}
}
