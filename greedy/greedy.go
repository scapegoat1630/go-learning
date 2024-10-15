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

// 134. 加油站
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
//给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
// 示例 1:
//
//输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
//输出: 3
//解释:
//从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
//开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
//开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
//开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
//开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
//开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
//因此，3 可为起始索引。
//示例 2:
//
//输入: gas = [2,3,4], cost = [3,4,3]
//输出: -1
//解释:
//你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
//我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
//开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
//开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
//你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
//因此，无论怎样，你都不可能绕环路行驶一周。
//
//
//提示:
//
//gas.length == n
//cost.length == n
//1 <= n <= 105
//0 <= gas[i], cost[i] <= 104

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < len(gas); {
		l := 0
		c := 0
		for {
			j := (i + c) % n
			l += gas[j] - cost[j]
			if l < 0 {
				break
			}
			if c == n {
				return i
			}
			c++
		}
		i = i + c
	}
	return -1
}
