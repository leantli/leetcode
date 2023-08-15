package main

// https://leetcode.cn/problems/min-cost-climbing-stairs/
// 746. 使用最小花费爬楼梯

func minCostClimbingStairs(cost []int) int {
	// 显然，很难去枚举回溯dfs各种情况，就算成功了时间复杂度也非常高
	// 这里显然考虑动态规划，定义 dp[i] 为到达台阶 i 时，最低花费多少
	// 则 dp[i] = min(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])，向上爬一楼便宜还是直接爬两楼便宜
	// 初始化，从下标 0 和 1 都可以开始，因此 dp[0] = dp[1] = 0
	// 当 len(cost) = 3 时，需要爬到第四层，也就是 dp[3]，因此 dp 长度为 len(cost)+1
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
