package union_find

// 并查集
/*
并查集（Union-find Data Structure）是一种树型的数据结构。它的特点是由子结点找
到父亲结点，用于处理一些不交集（Disjoint Sets）的合并及查询问题。
• Find：确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
• Union：将两个子集合并成同一个集合。
*/

func GetLongest(nums []int) int {
	m := make(map[int]bool)
	ans := 0
	for _, v := range nums {
		m[v] = true
	}
	for _, v := range nums {
		if m[v-1] {
			continue
		}
		t := 1
		for m[v+t] {
			t++
		}
		if t > ans {
			ans = t
		}
	}
	return ans
}

func GetRedundant(edges [][]int) []int {
	m := make([]int, len(edges)+1)
	for i := 0; i < len(m); i++ {
		m[i] = i
	}
	for _, v := range edges {
		x, y :=  find(m, v[0]), find(m, v[1])
		if find(m, x) != find(m, y) {
			m[x] = y
		} else {
			return []int{v[0], v[1]}
		}
	}
	return []int{}
}

func find(nums []int, n int) int {
	for nums[n] != n {
		n = nums[n]
	}
	return n
}