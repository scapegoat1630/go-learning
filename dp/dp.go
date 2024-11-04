package dp

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p, q = q, r
		r = p + q
	}
	return r
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	p, q := 1, 2
	for i := 3; i <= n; i++ {
		r := p + q
		if i == n {
			return r
		}
		p = q
		q = r
	}
	return -1
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// 1125
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n := len(req_skills)
	dp := make([][]int, 1<<n)
	skillsMap := make(map[string]int)
	for i, skill := range req_skills {
		skillsMap[skill] = 1 << i
	}
	dp[0] = []int{}
	for j, p := range people {
		curr := 0
		for _, skill := range p {
			curr |= skillsMap[skill]
		}
		for i := 0; i < (1 << n); i++ {
			if dp[i] == nil {
				continue
			}
			newMask := curr | i
			if dp[newMask] == nil || len(dp[i])+1 < len(dp[newMask]) {
				dp[newMask] = append([]int{j}, dp[i]...)
			}
		}
	}
	return dp[(1<<n)-1]
}

// 221. 最大正方形
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
// https://leetcode.cn/problems/maximal-square/description/?envType=study-plan-v2&envId=top-interview-150

func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				maxLen = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen * maxLen
}

// 97. 交错字符串
// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
//两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空
//子字符串
//：
//
//s = s1 + s2 + ... + sn
//t = t1 + t2 + ... + tm
//|n - m| <= 1
//交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//注意：a + b 意味着字符串 a 和 b 连接。

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	p := len(s3)
	if n+m != p {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			t := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[t]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[t]
			}
		}
	}
	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
