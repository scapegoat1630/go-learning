package two_point

import (
	"fmt"
	"math"
	"sort"
)

func MoveZeros(nums []int) {
	p, c := 0, 0
	for c < len(nums) {
		if nums[c] != 0 {
			nums[c], nums[p] = nums[p], nums[c]
			p++
		}
		c++
	}
}

func subSort(nums []int) []int {
	l := -1
	r := -1
	max := math.MinInt64
	min := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			l = i
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			r = i
		}
	}
	if l == r {
		return []int{-1, -1}
	}
	return []int{r, l}

}

func pairSums(nums []int, target int) [][]int {
	ans := [][]int{}
	m := map[int]int{}
	for i, num := range nums {
		if _, ok := m[target-num]; !ok {
			m[num] = i
		} else {
			ans = append(ans, []int{target - num, num})
			delete(m, target-num)
		}
	}
	return ans
}

func pairSum2(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] == target {
			ans = append(ans, []int{nums[l], nums[r]})
			l++
			r--
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			l++
		}
	}
	return ans
}

func maxArea(nums []int) int {
	l, r := 0, len(nums)-1
	max := math.MinInt64
	for l < r {
		if nums[l] <= nums[r] {
			area := (r - l) * nums[l]
			if area > max {
				max = area
			}
			l++
		} else {
			area := (r - l) * nums[r]
			if area > max {
				max = area
			}
			r--
		}
	}
	return max
}

func trap(nums []int) int {
	ans := 0
	leftMax := 0
	rightMax := 0
	l, r := 0, len(nums)-1
	for l <= r {
		if leftMax < rightMax {
			if nums[l] < leftMax {
				ans += leftMax - nums[l]
			} else {
				leftMax = nums[l]
			}
			l++
		} else {
			if nums[r] < rightMax {
				ans += rightMax - nums[r]
			} else {
				rightMax = nums[r]
			}
			r--
		}
	}
	return ans
}

func sortColors2(nums []int) {
	i := 0
	p1, p2 := 0, len(nums)-1
	for i <= p2 {
		if nums[i] == 0 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		i++
	}
}

func sortColors(nums []int) {
	p0, p1 := 0, len(nums)-1
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			fmt.Printf("%#v\n",nums)
		} else if c == 2 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1--
			fmt.Printf("%#v\n",nums)
		}
	}
}
