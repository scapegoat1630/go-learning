package search

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func GetMin(root *Node) int {
	f := false
	pre := math.MinInt32
	min := math.MaxInt32
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		dfs(n.Left)
		if f && min > n.Val-pre {
			min = n.Val - pre
		}
		f = true
		pre = n.Val
		dfs(n.Right)
	}
	dfs(root)
	return min
}

func trimBTS(n *Node, low, high int) *Node {
	if n == nil {
		return nil
	}
	if n.Val > high {
		return trimBTS(n.Left, low, high)
	}
	if n.Val < low {
		return trimBTS(n.Right, low, high)
	}
	n.Left = trimBTS(n.Left, low, high)
	n.Right = trimBTS(n.Right, low, high)
	return n
}

func averageOfLevels(n *Node) (averages []float64) {
	curr := []*Node{n}
	for len(curr) > 0 {
		sum := 0.0
		next := []*Node{}
		for _, v := range curr {
			if v != nil {
				sum += float64(v.Val)
				if v.Left != nil {
					next = append(next, v.Left)
				}
				if v.Right != nil {
					next = append(next, v.Right)
				}
			}
		}
		averages = append(averages, sum/float64(len(curr)))
		curr = next
	}
	return
}

func mySqrt(x int) int {
	l, r := 0, x
	mid := l + (r-l)/2
	for l <= r {
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return mid
}
