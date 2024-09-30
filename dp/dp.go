package dp

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p, q = q, r
		r = p + q
	}
	return r
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
